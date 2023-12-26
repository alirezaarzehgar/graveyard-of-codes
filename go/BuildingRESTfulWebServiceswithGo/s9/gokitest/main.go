package main

import (
	"gokitest/helpers"
	"log"
	"net/http"
	"os"

	httptransport "github.com/go-kit/kit/transport/http"
	logkit "github.com/go-kit/log"
)

func main() {
	svc := helpers.LoggingMiddleware{
		Logger: logkit.NewLogfmtLogger(os.Stderr),
		Next:   helpers.EncryptServiceInstance{},
	}

	http.Handle("/encrypt", httptransport.NewServer(helpers.MakeEncryptEndpoint(svc),
		helpers.DecodeEncriptRequest,
		helpers.EncodeResponse))
	http.Handle("/decrypt", httptransport.NewServer(helpers.MakeDecryptEndpoint(svc),
		helpers.DecodeDecryptRequest,
		helpers.EncodeResponse))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
