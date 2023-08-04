package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

func main() {
	data, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, data.Body)

	URL, _ := url.Parse("https://www.google.com")
	c := http.Client{
		Timeout: time.Second * 30,
	}
	req, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		log.Fatal("GET:", err)
	}
	httpData, err := c.Do(req)
	fmt.Println("status:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println(string(header))

	c = http.Client{
		Transport: &http.Transport{
			Dial: func(network, addr string) (net.Conn, error) {
				conn, _ := net.DialTimeout(network, addr, time.Second)
				conn.SetDeadline(time.Now().Add(time.Second))
				return conn, nil
			},
		},
	}

	data, err = c.Get(URL.String())
	if err != nil {
		log.Fatal(err)
	} else {
		defer data.Body.Close()
		io.Copy(os.Stdout, data.Body)
	}
}
