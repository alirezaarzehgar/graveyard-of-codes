package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need the template file!")
		return
	}

	entries := []struct{ Number, Square int }{
		{-1, 1}, {-2, 4}, {-3, 9}, {-4, 16},
	}

	t := template.Must(template.ParseGlob(os.Args[1]))
	t.Execute(os.Stdout, entries)

	t, err := template.New("Foo").Parse("{{ range . }} {{ printf \"%d vs %d\" .Number .Square }}\n{{ end }}")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, entries)

	t, err = template.New("Foo").Parse("{{3 -}} < {{- 4}}\n")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, entries)

	data := []int{}
	for i := 0; i < 120; i++ {
		data = append(data, i)
	}

	t, err = template.New("Foo").Parse("{{range .}}{{if .}} {{.}} {{end}}\n{{end}}")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, data)
}
