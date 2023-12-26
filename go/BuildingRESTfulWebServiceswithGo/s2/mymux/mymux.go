package mymux

import (
	"fmt"
	"net/http"
)

type mymux struct {
	routes map[string]http.HandlerFunc
}

func New() *mymux {
	mux := mymux{}
	mux.routes = make(map[string]http.HandlerFunc)
	return &mux
}

func (m *mymux) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "%s\n", req.URL.Path)

	handlers, ok := m.routes[req.URL.Path]
	if !ok {
		http.NotFound(w, req)
		return
	}
	handlers(w, req)
}

func (m *mymux) HandleFunc(route string, handler http.HandlerFunc) {
	m.routes[route] = handler
}
