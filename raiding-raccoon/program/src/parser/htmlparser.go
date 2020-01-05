package parser

import (
	"io"
	"net/url"

	"golang.org/x/net/html"
)

// HTMLParser struct.
// Implementation of the parser.Parser interface that parses
// data in the HTML format.
type HTMLParser struct {
}

var (
	// <a href="LINK">text</a>
	linkData = "a"
	linkKey  = "href"
)

// Parse func
func (htmlp *HTMLParser) Parse(reader io.ReadCloser, links chan<- *url.URL) {
	defer reader.Close()
	defer close(links)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()

		switch {
		case tokenType == html.ErrorToken: // End of html content
			return
		case tokenType == html.StartTagToken: // Start of a new tag
			t := tokenizer.Token()
			if t.Data == linkData {
				htmlp.extractLinkAttribute(&t, func(x *url.URL) { links <- x })
			}
		}
	}
}

func (htmlp *HTMLParser) extractLinkAttribute(token *html.Token, onLink func(*url.URL)) {
	for _, attr := range token.Attr {
		if attr.Key == linkKey {
			if uri, err := url.Parse(attr.Val); err == nil {
				onLink(uri)
			}
		}
	}
}
