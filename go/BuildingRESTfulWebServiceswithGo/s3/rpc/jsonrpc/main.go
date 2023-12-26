package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type Args struct {
	ID uint `json:"id"`
}

type Book struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Server struct{}

func (s *Server) GetBook(r *http.Request, args *Args, reply *Book) error {
	*reply = Book{
		ID:          args.ID,
		Name:        "Deep Work",
		Title:       "How to work deeply",
		Description: "A bullshit book",
	}
	return nil
}

func main() {
	s := rpc.NewServer()
	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterService(new(Server), "")

	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":8000", r)
}
