package crawler

import (
	"raiding-raccoon/src/crawler"
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestDomainSelector(t *testing.T) {
	sel := crawler.DomainSelector("test.com")
	match, _ := url.Parse("https://test.com:80/about")
	fail, _ := url.Parse("https://testx.com:80/about")
	assert.Assert(t, sel(match))
	assert.Assert(t, !sel(fail))
}
