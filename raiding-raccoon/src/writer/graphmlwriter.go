package writer

import (
	"io"
	"raiding-raccoon/src/graph"

	"github.com/shabbyrobe/xmlwriter"
)

// GraphMLWriter struct.
// Implementation of the writer.Writer interface that writes
// reference graphs in the GraphML format.
// See <http://graphml.graphdrawing.org/> for more information.
type GraphMLWriter struct {
}

// Write func
func (gmlw *GraphMLWriter) Write(graph graph.Graph, writer io.Writer) (err error) {
	// Don't pollute code with all the error checks. If something fails, the consistency
	// of the xml is not guaranteed anymore. Therefore return an error.
	defer func() {
		if xerr := recover(); xerr != nil {
			err = xerr.(error)
		}
	}()

	xmlw := xmlwriter.Open(writer)

	// writing the top level XML structure
	xmlw.StartDoc(xmlwriter.Doc{})
	rootAttrs := []xmlwriter.Attr{
		{
			Name:  "xmlns",
			Value: "http://graphml.graphdrawing.org/xmlns",
		},
		{
			Name:  "xmlns:xsi",
			Value: "http://www.w3.org/2001/XMLSchema-instance",
		},
		{
			Name:  "xsi:schemaLocation",
			Value: "http://graphml.graphdrawing.org/xmlns http://graphml.graphdrawing.org/xmlns/1.0/graphml.xsd",
		},
	}
	xmlw.StartElem(xmlwriter.Elem{
		Name:  "graphml",
		Attrs: rootAttrs,
	})
	// writes the reference graph into the xml writer
	gmlw.writeGraph(xmlw, graph)
	xmlw.EndElem("graphml")
	xmlw.EndAllFlush()

	return nil
}

func (gmlw *GraphMLWriter) writeGraph(writer *xmlwriter.Writer, source graph.Graph) error {
	attrs := []xmlwriter.Attr{
		{
			Name:  "id",
			Value: "graph",
		},
		{
			Name:  "edgedefault",
			Value: "directed",
		},
	}
	writer.StartElem(xmlwriter.Elem{
		Name:  "graph",
		Attrs: attrs,
	})

	gmlw.writeNodes(writer, source)
	gmlw.writeEdges(writer, source)

	writer.EndElem("graph")
	return nil
}

func (gmlw *GraphMLWriter) writeNodes(writer *xmlwriter.Writer, source graph.Graph) error {
	nodeChan := make(chan string)
	go source.Nodes(nodeChan)
	for node := range nodeChan {
		attrs := []xmlwriter.Attr{
			{
				Name:  "id",
				Value: node,
			},
		}
		writer.WriteElem(xmlwriter.Elem{
			Name:  "node",
			Attrs: attrs,
		})
	}
	return nil
}

func (gmlw *GraphMLWriter) writeEdges(writer *xmlwriter.Writer, source graph.Graph) error {
	edgeChan := make(chan graph.Edge)
	go source.Edges(edgeChan)
	for edge := range edgeChan {
		attrs := []xmlwriter.Attr{
			{
				Name:  "source",
				Value: edge.Source,
			},
			{
				Name:  "target",
				Value: edge.Dest,
			},
		}

		writer.WriteElem(xmlwriter.Elem{
			Name:  "edge",
			Attrs: attrs,
		})
	}
	return nil
}
