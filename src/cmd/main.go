// The main package is the console based starting point for the program.
package main

import (
	"net/url"
	"os"

	"raiding-raccoon/src/crawler"
	"raiding-raccoon/src/graph"
	"raiding-raccoon/src/loader"
	"raiding-raccoon/src/parser"
	"raiding-raccoon/src/writer"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

// main func
func main() {
	// parsing configuration parameters
	uri, err := configure()
	if err != nil {
		panic(err)
	}
	log.Infof("Using URI: \"%v\" as start", uri)

	// setting up the parts of the logic
	graph := graph.NewConcurrentMapGraph()
	c := crawler.NewCrawler(
		uri.Scheme, uri.Host,
		&loader.HTTPLoader{},
		&parser.HTMLParser{},
		graph)
	c.UseSelector(
		crawler.DomainSelector(crawler.RefineHostname(uri)),
		crawler.UniqueSelector(),
	)
	// enlisting first URI in order to start crawling and wait for it to finish
	c.Enlist(uri)
	c.Wait()

	// export the result graph into STDOUT
	writer := &writer.GraphMLWriter{}
	if err := writer.Write(graph, os.Stdout); err != nil {
		log.Errorf("Could not write graphml, %v", err)
	}
}

func configure() (*url.URL, error) {
	// print cwd, always good to have
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Infof("Current working directory: %v", wd)

	// define and parse flags
	var startURI string
	pflag.StringVar(&startURI, "start", "", "The start URI from which to go on crawling (including protocol)")
	pflag.Parse()

	// url.ParseRequestURI will need the protocol (http:// or https://)
	uri, err := url.ParseRequestURI(startURI)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return uri, nil
}
