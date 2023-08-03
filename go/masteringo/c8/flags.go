package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	names []string
}

func (n NamesFlag) GetNames() []string {
	return n.names
}

func (n *NamesFlag) String() string {
	return fmt.Sprint(n.names)
}

func (n *NamesFlag) Set(v string) error {
	if len(n.names) > 0 {
		return errors.New("Cannot use names flag more that once!")
	}

	names := strings.Split(v, ",")

	for _, name := range names {
		n.names = append(n.names, name)
	}
	return nil
}

func main() {
	var manyNames NamesFlag
	n := flag.Int("number", 0, "set a number")
	flag.IntVar(n, "n", 0, "set a number")
	flag.Var(&manyNames, "names", "Get list of names")
	flag.Parse()

	fmt.Println("number:", *n)
	fmt.Println("List of names:")
	for _, name := range manyNames.GetNames() {
		fmt.Println(name)
	}
}
