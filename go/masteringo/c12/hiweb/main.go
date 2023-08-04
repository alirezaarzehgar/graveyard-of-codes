package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/mbndr/figlet4go"
)

type HandlerType func() map[string]any

type Handler struct {
	HandlerFuncs map[string]HandlerType
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%#v\n", r.URL)

	var data map[string]any
	pattern := strings.Split(r.URL.String(), "?")[0]
	handler, ok := h.HandlerFuncs[pattern]
	if ok {
		data = handler()
	} else {
		data = h.NotFound()
	}
	statusCode, ok := data["status"].(int)
	if !ok {
		statusCode = http.StatusOK
	}
	delete(data, "status")

	out, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(out)
}

func (h *Handler) NotFound() map[string]any {
	return map[string]any{
		"message": "Not found!",
		"status":  http.StatusNotFound,
	}
}

func Root() map[string]any {
	return map[string]any{
		"message": "Fucking root say hello!",
		"status":  http.StatusOK,
	}
}

func Time() map[string]any {
	return map[string]any{
		"message": "Time:" + time.Now().Format(time.RFC1123),
	}
}

func main() {
	ascii := figlet4go.NewAsciiRender()
	out, _ := ascii.Render("Hello, Go!")
	fmt.Println(out)

	handler := &Handler{map[string]HandlerType{
		"/":     Root,
		"/time": Time,
	}}

	err := http.ListenAndServe(":8000", handler)
	if err != nil {
		log.Fatal(err)
	}
}
