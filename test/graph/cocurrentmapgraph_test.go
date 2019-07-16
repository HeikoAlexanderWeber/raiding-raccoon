package graph

import (
	"fmt"
	"raiding-raccoon/src/graph"
	"testing"

	"gotest.tools/assert"
)

func TestAddNode(t *testing.T) {
	graph := graph.NewConcurrentMapGraph()
	// Unique items are ok
	assert.Assert(t, graph.AddNode("a"))
	assert.Assert(t, graph.AddNode("b"))
	// Duplicate items are rejected
	assert.Assert(t, !graph.AddNode("a"))
}

func TestAddEdge(t *testing.T) {
	graph := graph.NewConcurrentMapGraph()
	assert.Assert(t, graph.AddEdge("a", "b"))
	assert.Assert(t, !graph.AddEdge("a", "b"))
	assert.Assert(t, graph.AddEdge("a", "c"))
	assert.Assert(t, graph.AddEdge("b", "a"))
}

func TestAddNodeAndCheckNodes(t *testing.T) {
	graph := graph.NewConcurrentMapGraph()
	assert.Assert(t, graph.AddNode("a"))
	assert.Assert(t, graph.AddNode("b"))
	assert.Assert(t, !graph.AddNode("a"))

	nodes := make(chan string)
	nodesdict := make(map[string]byte)
	go graph.Nodes(nodes)
	for n := range nodes {
		_, exists := nodesdict[n]
		assert.Assert(t, !exists)
		nodesdict[n] = 0
	}

	for k := range nodesdict {
		assert.Assert(t, k == "a" || k == "b")
	}
	assert.Assert(t, len(nodesdict) == 2)
}

func TestAddEdgeAddsNodes(t *testing.T) {
	graph := graph.NewConcurrentMapGraph()
	assert.Assert(t, graph.AddEdge("a", "b"))
	assert.Assert(t, !graph.AddEdge("a", "b"))
	assert.Assert(t, graph.AddEdge("a", "c"))
	assert.Assert(t, graph.AddEdge("b", "a"))

	nodes := make(chan string)
	nodesdict := make(map[string]byte)
	go graph.Nodes(nodes)
	for n := range nodes {
		_, exists := nodesdict[n]
		assert.Assert(t, !exists)
		nodesdict[n] = 0
	}

	for k := range nodesdict {
		assert.Assert(t, k == "a" || k == "b" || k == "c")
	}
	assert.Assert(t, len(nodesdict) == 3)
}

func TestAddEdgeAndCheckEdges(t *testing.T) {
	g := graph.NewConcurrentMapGraph()
	assert.Assert(t, g.AddEdge("a", "b"))
	assert.Assert(t, g.AddEdge("a", "c"))
	assert.Assert(t, g.AddEdge("b", "a"))

	edges := make(chan graph.Edge)
	nodesdict := make(map[string]byte)
	go g.Edges(edges)
	for n := range edges {
		key := fmt.Sprintf("%v->%v", n.Source, n.Dest)
		_, exists := nodesdict[key]
		assert.Assert(t, !exists)
		nodesdict[key] = 0
	}
	for k := range nodesdict {
		assert.Assert(t,
			k == "a->b" || k == "a->c" || k == "b->a")
	}
	assert.Assert(t, len(nodesdict) == 3)
}

func TestIterateCb(t *testing.T) {
	g := graph.NewConcurrentMapGraph()
	assert.Assert(t, g.AddEdge("a", "b"))
	assert.Assert(t, g.AddEdge("b", "c"))
	assert.Assert(t, g.AddEdge("c", "a"))

	edgeDict := make(map[string]byte)
	g.IterateCb(func(x string, y string) {
		edgeDict[fmt.Sprintf("%v->%v", x, y)] = 0
	})
	assert.Assert(t, len(edgeDict) == 3)
	for k := range edgeDict {
		assert.Assert(t,
			k == "a->b" || k == "b->c" || k == "c->a")
	}
}
