package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("This is"))
	fmt.Fprintf(&buf, " a string!\n")
	buf.WriteTo(os.Stdout)
	buf.WriteTo(os.Stdout)
}
