package graph

import (
	"fmt"

	cmap "github.com/orcaman/concurrent-map"
	log "github.com/sirupsen/logrus"
)

// ConcurrentMapGraph struct.
// Implementation of the graph.Graph interface using
// a concurrent map strategy. This type is implemented
// for safe concurrent use.
type ConcurrentMapGraph struct {
	data cmap.ConcurrentMap
}

// NewConcurrentMapGraph func
func NewConcurrentMapGraph() *ConcurrentMapGraph {
	return &ConcurrentMapGraph{
		data: cmap.New(),
	}
}

// AddNode func
func (graph *ConcurrentMapGraph) AddNode(id string) bool {
	if graph.data.SetIfAbsent(id, cmap.New()) {
		log.Infof("Added %v to graph", id)
		return true
	}
	return false
}

// AddEdge func
func (graph *ConcurrentMapGraph) AddEdge(source string, dest string) bool {
	// Ensure that src and dest nodes exist in the graph
	graph.AddNode(source)
	graph.AddNode(dest)

	// Add the reference from src --> dest
	if entry, exists := graph.data.Get(source); exists {
		if entry.(cmap.ConcurrentMap).SetIfAbsent(dest, byte(0)) {
			log.Infof("Added edge %v --> %v", source, dest)
			return true
		}
		return false
	}
	err := fmt.Errorf("Invalid state, entry was not found (source: \"%v\", dest: \"%v\")", source, dest)
	log.Error(err)
	panic(err) // logic error
}

// Nodes func
func (graph *ConcurrentMapGraph) Nodes(out chan<- string) {
	defer close(out)
	graph.data.IterCb(func(src string, _ interface{}) {
		out <- src
	})
}

// Edges func
func (graph *ConcurrentMapGraph) Edges(out chan<- Edge) {
	defer close(out)
	graph.IterateCb(func(x string, y string) {
		out <- Edge{
			Source: x,
			Dest:   y,
		}
	})
}

// IterateCb func
func (graph *ConcurrentMapGraph) IterateCb(cb func(string, string)) {
	graph.data.IterCb(func(src string, v interface{}) {
		subMap := v.(cmap.ConcurrentMap)
		subMap.IterCb(func(dest string, _ interface{}) {
			cb(src, dest)
		})
	})
}
