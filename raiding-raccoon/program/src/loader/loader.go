package loader

import (
	"io"
	"net/url"
)

// Loader interface.
// Used for loading data.
type Loader interface {
	// Loads data from the given URL.
	Load(*url.URL) (io.ReadCloser, error)
}
