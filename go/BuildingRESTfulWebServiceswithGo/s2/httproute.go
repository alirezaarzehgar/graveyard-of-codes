package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(name string, args ...string) string {
	cmd := exec.Command(name, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return stdout.String()
}

func goVersion(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, getCommandOutput("go", "version"))
}

func getFileContent(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprint(w, getCommandOutput("cat", p.ByName("name")))
}

func main() {
	router := httprouter.New()
	router.GET("/api/v1/go-version", goVersion)
	router.GET("/api/v1/show-file/:name", getFileContent)
	log.Fatal(http.ListenAndServe(":8000", router))
}
