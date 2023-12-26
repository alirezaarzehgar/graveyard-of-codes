package main

import (
	"io"
	"net/http"

	"github.com/emicklei/go-restful"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/ping").To(Pong))
	restful.Add(ws)
	http.ListenAndServe(":8000", nil)
}

func Pong(r *restful.Request, w *restful.Response) {
	io.WriteString(w, "pong")
}
