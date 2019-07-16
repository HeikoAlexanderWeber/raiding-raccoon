package graph

// Graph interface.
// Interface for a string to string reference graph.
type Graph interface {
	// Will add a node to the graph if not already existing.
	// Returns whether the process was successful.
	AddNode(string) bool
	// Will add a reference between two nodes if it does not
	// already exist. Nodes that do not exist yet will be created.
	// Returns whether the reference was added successfully.
	AddEdge(string, string) bool
	// Will write all the nodes in the graph to the given channel.
	Nodes(chan<- string)
	// Will write all the edges in the graph to the given channel.
	Edges(chan<- Edge)
	// Will iterate through every reference between nodes with src, dest.
	IterateCb(func(string, string))
}

// Edge struct.
// Describing an edge between two nodes in a graph.
type Edge struct {
	Source string
	Dest   string
}
