package crawler

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestUniqueSelector(t *testing.T) {
	sel := UniqueSelector()
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
