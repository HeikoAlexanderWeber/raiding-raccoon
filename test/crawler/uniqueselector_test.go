package crawler

import (
	"net/url"
	"raiding-raccoon/src/crawler"
	"testing"

	"gotest.tools/assert"
)

func TestUniqueSelector(t *testing.T) {
	sel := crawler.UniqueSelector()
	a, _ := url.Parse("https://test.com:80/about")
	b, _ := url.Parse("https://testx.com:80/about")
	c, _ := url.Parse("https://test.com:80/1")
	assert.Assert(t, sel(a))
	assert.Assert(t, sel(b))
	assert.Assert(t, sel(c))
	assert.Assert(t, !sel(a))
	assert.Assert(t, !sel(b))
	assert.Assert(t, !sel(c))
}
