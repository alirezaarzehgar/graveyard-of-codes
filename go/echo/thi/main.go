package main

import (
	"html/template"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	t, _ := template.New("hi").Parse(`
	<h1>{{.N1}}</h1>
	<p>{{.N2}}</p>

	{{range .Items}} <div>{{.}}</div> {{else}} <strong>no rows</strong> {{end}}
	`)
	t.Execute(w, struct {
		N1    string
		N2    string
		Items []string
	}{N1: "hello", N2: "hi", Items: []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}})
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
