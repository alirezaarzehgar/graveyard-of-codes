package main

import (
	"io"
	"net/http"

	"github.com/emicklei/go-restful"
)

func Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/api").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("res").To(func(r *restful.Request, w *restful.Response) { io.WriteString(w, "get") }))
	ws.Route(ws.POST("res").To(func(r *restful.Request, w *restful.Response) { io.WriteString(w, "post") }))
	ws.Route(ws.PUT("res").To(func(r *restful.Request, w *restful.Response) { io.WriteString(w, "put") }))
	ws.Route(ws.DELETE("res").To(func(r *restful.Request, w *restful.Response) { io.WriteString(w, "delete") }))

	container.Add(ws)
}

func main() {
	container := restful.NewContainer()
	container.Router(restful.CurlyRouter{})
	Register(container)
	s := http.Server{Addr: ":8000", Handler: container}
	s.ListenAndServe()
}
