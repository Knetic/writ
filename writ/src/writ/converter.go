package writ

import (
	"strings"
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

		// strip out carriage returns, this renderer doesn't handle them well.
		sanitized := fixLineEndings(string(content))

		req.output = markdown.ToHTML([]byte(sanitized), nil, nil)
		req.done <- nil
	}
}

func fixLineEndings(in string) string {
	in = strings.Replace(in, "\r\n", "\n", -1)
	in = strings.Replace(in, "\r", "\n", -1)
	return in;
}