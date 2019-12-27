package writ

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"encoding/json"
)

type Server struct {
	converter chan *convertRequest
}

func NewServer() *Server {
	return &Server{}
}

func (this *Server) Listen(path string) error {

	this.converter = make(chan *convertRequest)
	go runConverter(this.converter)

	mux := http.NewServeMux()
	mux.HandleFunc("/", this.serve)
	mux.HandleFunc("/list", this.list)

	return http.ListenAndServe(path, mux)
}

func (this Server) serve(response http.ResponseWriter, request *http.Request) {

	path := "." + request.URL.Path + ".md"
	body, err := this.convertFile(path)
	if err != nil {
		writeError(response, err)
		return
	}

	response.Write(body)
}

// Returns an array of all md files in or underneath the cwd
func (this Server) list(response http.ResponseWriter, request *http.Request) {

	items, _ := filepath.Glob("./**.md")
	for i := 0; i < len(items); i++ {
		
		item := items[i]
		if strings.HasSuffix(item, ".md") {
			item = item[:len(item)-3]
			items[i] = item
		}
	}

	r := ListFiles {
		Items: items,
	}

	serialized, err := json.Marshal(r)
	if err != nil {
		writeError(response, err)
		return
	}
	
	response.Write(serialized)
}

func (this Server) convertFile(path string) ([]byte, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	done := make(chan error)

	req := convertRequest {
		reader: f,
		done: done,
	}
	this.converter <- &req
	err = <-done

	return req.output, err
}

func writeError(response http.ResponseWriter, err error) {
	response.WriteHeader(500)
	response.Write([]byte(err.Error()))		
}

type ListFiles struct {
	Items []string `json:"items"`
}