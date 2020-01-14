package crawler

import (
	"net/url"
	"testing"

	cmap "github.com/orcaman/concurrent-map"
	"gotest.tools/assert"
)

func TestUniqueSelector(t *testing.T) {
	uniqueMap := cmap.New()
	sel := UniqueSelector(
		func(d string) bool {
			return uniqueMap.SetIfAbsent(d, byte(0))
		})
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
