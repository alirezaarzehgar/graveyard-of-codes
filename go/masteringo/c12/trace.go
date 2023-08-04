package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptrace"
	"os"
)

func main() {
	c := http.Client{}
	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	trace := &httptrace.ClientTrace{
		GetConn: func(hostPort string) {
			fmt.Println("GetConn")
		},
		GotConn: func(gci httptrace.GotConnInfo) {
			fmt.Println("GotConn")
		},
		GotFirstResponseByte: func() {
			fmt.Println("GotFirstResponseByte")
		},
		DNSStart: func(di httptrace.DNSStartInfo) {
			fmt.Println("DNSStart")
		},
		DNSDone: func(di httptrace.DNSDoneInfo) {
			fmt.Println("DNSStart")
		},
		ConnectStart: func(network, addr string) {
			fmt.Println("ConnectStart")
		},
		ConnectDone: func(network, addr string, err error) {
			fmt.Println("ConnectDone")
		},
	}

	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))
	fmt.Println("Requesting data from server!")
	_, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		log.Fatal(err)
	}

	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, res.Body)
}
