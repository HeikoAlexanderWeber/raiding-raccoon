package writer

import (
	"io"
	"raiding-raccoon/program/src/graph"
)

// Writer interface.
// Used for exporting a reference graph into some format.
type Writer interface {
	// Writes the given reference graph into the given io.Writer.
	// Returns any error that possibly occurred.
	Write(graph.Graph, io.Writer) error
}
