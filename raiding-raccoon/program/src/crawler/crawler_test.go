package crawler

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"raiding-raccoon/program/src/graph"
	"sort"
	"testing"

	"gotest.tools/assert"
)

func TestCrawler(t *testing.T) {
	g := graph.NewConcurrentMapGraph()
	c := NewCrawler(
		"", "",
		&loaderMock{},
		&parserMock{},
		g,
	)

	// parserMock always returns the same mock URI which would
	// lead to an infinite loop
	c.UseSelector(UniqueSelector())

	uri0, _ := url.Parse("https://localhost:80/about")
	c.Enlist(uri0)
	c.Wait()

	edges := []string{}
	g.IterateCb(func(x string, y string) {
		edges = append(edges, fmt.Sprintf("%v --> %v", x, y))
	})
	assert.Assert(t, len(edges) == 2)
	sort.Strings(edges)
	assert.Equal(t, edges[0], "https://localhost:80/about --> https://localhost:80/one")
	assert.Equal(t, edges[1], "https://localhost:80/one --> https://localhost:80/one")
}

type loaderMock struct {
}

func (l *loaderMock) Load(uri *url.URL) (io.ReadCloser, error) {
	return ioutil.NopCloser(bytes.NewReader([]byte(""))), nil
}

type parserMock struct {
}

func (p *parserMock) Parse(rc io.ReadCloser, uris chan<- *url.URL) {
	defer rc.Close()
	defer close(uris)
	uri1, _ := url.Parse("https://localhost:80/one")
	uris <- uri1
}
