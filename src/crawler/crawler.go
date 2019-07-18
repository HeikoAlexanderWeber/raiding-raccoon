package crawler

import (
	"net/url"
	"raiding-raccoon/src/graph"
	"raiding-raccoon/src/loader"
	"raiding-raccoon/src/parser"
	"sync"

	log "github.com/sirupsen/logrus"
)

// Crawler struct.
// Contains logic to glue loader, parser and graph builder together.
type Crawler struct {
	baseProtocol string
	baseDomain   string
	loader       loader.Loader
	parser       parser.Parser
	graph        graph.Graph
	selectors    []Selector
	parsersWg    sync.WaitGroup
}

// NewCrawler func.
// Returns a reference to a new crawler object.
func NewCrawler(baseProtocol string, baseDomain string,
	loader loader.Loader, parser parser.Parser, graph graph.Graph) *Crawler {
	return &Crawler{
		baseProtocol: baseProtocol,
		baseDomain:   baseDomain,
		loader:       loader,
		parser:       parser,
		graph:        graph,
		selectors:    []Selector{},
		parsersWg:    sync.WaitGroup{},
	}
}

// Enlist func.
// This function is going to deeply crawl the URI that is given as parameter
// if the selectors are approving the given URI.
func (crawler *Crawler) Enlist(uri *url.URL) {
	// block one crawler process
	crawler.parsersWg.Add(1)
	// invoke the crawler logic
	go func() {
		// if it exits, always mark the process as finished
		defer crawler.parsersWg.Done()
		// get selectors approval
		if !crawler.filter(uri) {
			return
		}
		// in every case, this will be a node
		crawler.graph.AddNode(uri.String())
		log.Infof("Enlisted: %v", uri.String())

		// load data from given URI
		reader, err := crawler.loader.Load(uri)
		if err != nil {
			log.Error(err)
			return
		}

		// parse all the links from the loaded data
		links := make(chan *url.URL)
		go crawler.parser.Parse(reader, links)
		crawler.handleNewLinks(uri, links)
	}()
}

func (crawler *Crawler) handleNewLinks(uri *url.URL, data <-chan *url.URL) {
	for link := range data {
		newLink := *link
		if newLink.Scheme == "" { // relative URL
			newLink.Scheme = crawler.baseProtocol
		}
		if newLink.Host == "" { // relative URL
			newLink.Host = crawler.baseDomain
		}
		// enlist all the new links
		crawler.Enlist(&newLink)
		// add the src->dest reference to the reference graph
		crawler.graph.AddEdge(uri.String(), newLink.String())
	}
}

// Wait func.
// Wait for all the recursive crawling processes to finish.
func (crawler *Crawler) Wait() {
	crawler.parsersWg.Wait()
}

// UseSelector func.
// Append a selector as kind of middleware in order to filter valid URIs.
func (crawler *Crawler) UseSelector(fn ...Selector) {
	crawler.selectors = append(crawler.selectors, fn...)
}

func (crawler *Crawler) filter(url *url.URL) bool {
	for _, selector := range crawler.selectors {
		if !selector(url) {
			return false
		}
	}
	return true
}
