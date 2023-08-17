package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const NS = "ali"

var counter = prometheus.NewCounter(prometheus.CounterOpts{
	Namespace: NS,
	Name:      "my_counter",
	Help:      "yoyoyoyoyo",
})

var gauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: NS,
	Name:      "my_gauge",
	Help:      "yoyoyoyoyoyoyo",
})

var histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
	Namespace: NS,
	Name:      "my_histogram",
	Help:      "yoyoyoyoyoyoyoyoyo",
})

var summary = prometheus.NewSummary(prometheus.SummaryOpts{
	Namespace: NS,
	Name:      "my_summary",
	Help:      "yoyoyoyoyoyoyo",
})

func main() {
	prometheus.MustRegister(counter)
	prometheus.MustRegister(gauge)
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(summary)

	go func() {
		for {
			counter.Add(rand.Float64() * 5)
			gauge.Add(rand.Float64()*15 - 5)
			histogram.Observe(rand.Float64() * 10)
			summary.Observe(rand.Float64() * 10)
			time.Sleep(time.Second * 2)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)
}
