package parser

import (
	"io"
	"net/url"
)

// Parser interface.
// Used for parsing data from raw data.
type Parser interface {
	// Parses all the URIs from the given data.
	Parse(io.ReadCloser, chan<- *url.URL)
}
