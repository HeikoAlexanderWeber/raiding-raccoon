package loader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

// HTTPLoader struct
type HTTPLoader struct {
}

// Load func
func (httpl *HTTPLoader) Load(uri *url.URL) (io.ReadCloser, error) {
	resp, err := http.Get(uri.String())
	if err != nil {
		log.Errorf("%v : %v", uri.String(), err)
		return nil, err
	}
	if resp.StatusCode == 200 {
		log.Infof("Loaded: %v", uri.String())
		return resp.Body, nil
	}
	err = fmt.Errorf("GET on %v returned %v", uri.String(), resp.Status)
	log.Error(err)
	return nil, err
}
