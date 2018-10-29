package server

import (
	"net/http"

	"github.com/google/go-github/v18/github"

)

type HTTPServer struct {
	CAFile   string
	CertFile string
	KeyFile  string
	Addr string
}

func (h *HTTPServer) handler(w http.ResponseWriter, r *http.Request){
	client := github.NewClient()
}

func (h *HTTPServer) ListenAndServer() {
	http.HandleFunc("",h.handler)
	http.ListenAndServe(h.Addr,nil)
}

