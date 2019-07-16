package parser

import (
	"bytes"
	"io/ioutil"
	"raiding-raccoon/src/parser"
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestParseLinks(t *testing.T) {
	p := parser.HTMLParser{}
	sampleData, _ := ioutil.ReadFile("sample_html.html")
	reader := ioutil.NopCloser(bytes.NewReader(sampleData))

	links := make(chan *url.URL)
	allLinks := []*url.URL{}
	go p.Parse(reader, links)
	for l := range links {
		allLinks = append(allLinks, l)
	}
	assert.Assert(t, len(allLinks) == 1)
	assert.Equal(t, allLinks[0].String(), "https://test.com/test123")
}
