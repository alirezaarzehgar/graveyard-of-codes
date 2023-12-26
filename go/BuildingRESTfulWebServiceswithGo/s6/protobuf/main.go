package main

import (
	"encoding/json"
	"fmt"
	"protobuf/protofiles"

	"google.golang.org/protobuf/proto"
)

func main() {
	person := &protofiles.Person{
		Name:     "Alireza",
		Age:      15,
		UniScore: float64(12.3),
		Phones: []*protofiles.Person_PhoneNumber{
			{Number: "0123123", Type: protofiles.Person_HOME},
			{Number: "0123346", Type: protofiles.Person_WORK},
			{Number: "23123124", Type: protofiles.Person_MOBILE},
		},
	}

	body, _ := proto.Marshal(person)
	jsonBody, _ := json.Marshal(person)
	fmt.Printf("%+v", string(jsonBody))
	var p1 protofiles.Person
	proto.Unmarshal(body, &p1)
	fmt.Printf("%+v", p1)
}
