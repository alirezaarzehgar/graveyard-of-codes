package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// PORT is the TCP port number the server will listen to
var PORT = ":1234"

var g1 = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_goroutines",
		Help:      "Number of goroutines"})

var g2 = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Namespace: "packt",
		Name:      "n_memory",
		Help:      "Memory usage"})

func main() {
	prometheus.MustRegister(g1)
	prometheus.MustRegister(g2)

	http.Handle("/metrics", promhttp.Handler())

	go func() {
		for {
			time.Sleep(time.Second)
			g1.Set(10)
			g2.Set(10)
		}
	}()

	log.Println("Listening to port", PORT)
	log.Println(http.ListenAndServe(PORT, nil))
}
