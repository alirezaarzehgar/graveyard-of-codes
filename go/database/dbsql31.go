package main

import (
	"database/sql"
	"errors"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pterm/pterm"
)

type Person struct {
	id   uint
	name string
	age  uint64
}

func searchPerson(db *sql.DB) (*Person, error) {
	key, _ := pterm.DefaultInteractiveContinue.
		WithOptions([]string{"id", "name", "age"}).
		Show("by")

	value, _ := pterm.DefaultInteractiveTextInput.Show("Enter " + key)

	flsearch := make(map[string]Person)
	var titles []string

	q, err := db.Query("SELECT id, name, age FROM person WHERE "+key+" = ?", value)
	if err != nil {
		return nil, err
	}

	for q.Next() {
		var p Person
		err := q.Scan(&p.id, &p.name, &p.age)
		if err != nil {
			pterm.Error.Println(err)
		}

		title := pterm.Sprintf("id: %d, name: %s, age: %d", p.id, p.name, p.age)
		titles = append(titles, title)
		flsearch[title] = p
	}

	var p Person
	if titles == nil {
		return nil, errors.New("User not found")
	} else if len(titles) == 1 {
		for _, v := range flsearch {
			p = v
		}
	} else {
		title, _ := pterm.DefaultInteractiveSelect.
			WithOptions(titles).
			Show()
		p = flsearch[title]
	}

	return &Person{id: p.id, name: p.name, age: p.age}, nil
}

func main() {
	db, err := sql.Open("sqlite3", "dbsql31.db")
	if err != nil {
		pterm.Error.Println(err)
	}
	defer db.Close()

	// Migrate database
	stmt, err := db.Prepare(
		"CREATE TABLE IF NOT EXISTS person (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, age INTEGER);")
	if err != nil {
		pterm.Error.Println(err)
	}
	stmt.Exec()

	opt := ""
	for {
		opt, _ = pterm.DefaultInteractiveSelect.
			WithDefaultOption(opt).
			WithOptions([]string{pterm.Red("Exit"), "Show people", "Add person", "Update person", "Delete person"}).
			Show()

		switch opt {
		case pterm.Red("Exit"):
			os.Exit(0)
		case "Show people":
			rows, err := db.Query("SELECT id, name, age FROM person")
			if err != nil {
				pterm.Error.Println(err)
				break
			}

			persons := make(map[string]Person)
			for rows.Next() {
				var p Person
				err := rows.Scan(&p.id, &p.name, &p.age)
				if err != nil {
					pterm.Error.Println(err)
					continue
				}

				persons[pterm.Sprintf("%d:%s", p.id, p.name)] = p
			}

			var titles []string = []string{pterm.Red("Exit")}
			for v := range persons {
				titles = append(titles, v)
			}

			name := ""
			for {
				name, _ = pterm.DefaultInteractiveSelect.
					WithOptions(titles).
					WithDefaultOption(name).
					Show()

				if name == pterm.Red("Exit") {
					break
				}

				p := persons[name]
				pterm.Info.Printfln("id: %d, name: %s, age: %d", p.id, p.name, p.age)
			}
		case "Add person":
			p := Person{}
			p.name, _ = pterm.DefaultInteractiveTextInput.Show("Enter name")
			age, _ := pterm.DefaultInteractiveTextInput.Show("Enter age")
			p.age, err = strconv.ParseUint(age, 10, 64)
			if err != nil {
				pterm.Error.Println(err)
			}

			r, err := db.Prepare("INSERT INTO person (name, age) VALUES (?, ?)")
			if err != nil {
				pterm.Error.Println(err)
				break
			}

			_, err = r.Exec(p.name, p.age)
			if err != nil {
				pterm.Error.Println(err)
				break
			}

			pterm.Success.Printfln("Person %s created successfully", p.name)

		case "Update person":
			p, err := searchPerson(db)
			if err != nil {
				pterm.Error.Println(err)
				break
			}

			name, _ := pterm.DefaultInteractiveTextInput.Show("Edit name")
			age, _ := pterm.DefaultInteractiveTextInput.Show("Edit age")

			stmt, _ := db.Prepare("UPDATE person SET name = ?, age = ? WHERE id = ?")
			_, err = stmt.Exec(name, age, p.id)
			if err != nil {
				pterm.Error.Println(err)
			}

		case "Delete person":
			p, err := searchPerson(db)
			if err != nil {
				pterm.Error.Println(err)
				break
			}

			stmt, _ := db.Prepare("DELETE FROM person WHERE id = ?")
			pterm.Printfln("%#v", p)
			_, err = stmt.Exec(p.id)
			if err != nil {
				pterm.Error.Println(err)
			}
		}
	}
}
