package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Record struct {
	Firstname   string
	Lastname    string
	Phonenumber string
	Gender      string
	Online      bool

	Contacts []struct {
		Name        string
		Phonenumber string
	}
}

func main() {
	data := Record{Firstname: "Alireza", Lastname: "Arzehgar",
		Phonenumber: "09155093114", Gender: "Male", Online: true,
		Contacts: []struct {
			Name        string
			Phonenumber string
		}{
			{Name: "Mohammad", Phonenumber: "1231234123"},
			{Name: "Hosein", Phonenumber: "5544345"},
			{Name: "Ahmad", Phonenumber: "87564576"},
			{Name: "Reza", Phonenumber: "885943454"},
		}}

	xmlData, err := xml.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(xml.Header + string(xmlData))
}
