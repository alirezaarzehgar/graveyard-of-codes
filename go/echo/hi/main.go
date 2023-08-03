package main

import (
	"github.com/alirezaarzehgar/hi/route"
)

func main() {
	router := route.Init()
	router.Logger.Fatal(router.Start(":8080"))
}
