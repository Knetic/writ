package writ

import (
	"net/http"
	"os"
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

	return http.ListenAndServe(path, mux)
}

func (this Server) serve(response http.ResponseWriter, request *http.Request) {

	path := "." + request.URL.Path + ".md"
	body, err := this.convertFile(path)
	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))		
		return
	}

	response.Write(body)
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