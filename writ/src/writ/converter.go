package writ

import (
	"io"
	"io/ioutil"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
)

// Encapsulates a request to convert the given `reader`'s contents from md to html
// and write them to the given `writer`, signalling on `done` when complete.
type convertRequest struct {
	reader io.Reader
	output []byte
	done chan error
}

// (goroutine) indefinitely accepts requests through `in`, and converts them.
func runConverter(in chan *convertRequest) {

	// don't let browsers get any ideas about using style quotes.
	html.Escaper['"'] = []byte("&#34;")

	for req := range in {

		content, err := ioutil.ReadAll(req.reader)
		if err != nil {
			req.done <- err
			continue
		}
		
		req.output = markdown.ToHTML(content, nil, nil)
		req.done <- nil
	}
}