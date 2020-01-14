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

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// main func
func main() {
	f, err := os.OpenFile("./log", os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.StandardLogger().SetOutput(f)

	// parsing configuration parameters
	config, err := configure()
	if err != nil {
		panic(err)
	}
	log.Infof("Using URI: \"%v\" as start", config.StartURI)

	// setting up the parts of the logic
	informationIsCached := false
	var g graph.Graph
	if config.useRedis() {
		redisGraph := graph.NewRedisGraph(config.StartURI.String(), true)
		informationIsCached = redisGraph.Exists(config.StartURI.String(), true)
		g = redisGraph
	} else {
		g = graph.NewConcurrentMapGraph()
	}

	if !informationIsCached {
		c := crawler.NewCrawler(
			config.StartURI.Scheme, config.StartURI.Host,
			&loader.HTTPLoader{},
			&parser.HTMLParser{},
			g)

		var uniqueMap graph.Writer
		if config.useRedis() {
			uid, _ := uuid.NewRandom()
			uniqueMap = graph.NewRedisGraph(uid.String(), true)
		} else {
			uniqueMap = graph.NewConcurrentMapGraph()
		}
		c.UseSelector(
			crawler.DomainSelector(crawler.RefineHostname(config.StartURI)),
			crawler.UniqueSelector(
				func(d string) bool {
					return uniqueMap.AddNode(d)
				}),
		)
		// enlisting first URI in order to start crawling and wait for it to finish
		c.Enlist(config.StartURI)
		c.Wait()
	}

	// export the result graph into STDOUT
	writer := &writer.GraphMLWriter{}
	if err := writer.Write(g, os.Stdout); err != nil {
		log.Errorf("Could not write graphml, %v", err)
	}
}

type config struct {
	StartURI      *url.URL
	RedisBackbone string
}

func (c *config) useRedis() bool {
	return c.RedisBackbone != ""
}

func configure() (*config, error) {
	// print cwd, always good to have
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Infof("Current working directory: %v", wd)

	// url.ParseRequestURI will need the protocol (http:// or https://)
	startURI := os.Getenv("RR_START_URL")
	uri, err := url.ParseRequestURI(startURI)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	redisBackbone := os.Getenv("RR_REDIS_BACKBONE")

	return &config{
		StartURI:      uri,
		RedisBackbone: redisBackbone,
	}, nil
}
