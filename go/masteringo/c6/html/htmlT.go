package main

import (
	"html/template"
	"net/http"
)

type Entry struct {
	Number, Double, Square int
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	data := []Entry{}
	for i := 0; i < 20; i++ {
		data = append(data, Entry{i, i * 2, i * i})
	}

	t := template.Must(template.ParseGlob("template.gohtml"))
	t.ExecuteTemplate(w, "template.gohtml", data)

}

func main() {
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":8080", nil)
}
