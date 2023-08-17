package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var visitor = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace: "ali",
	Name:      "visit_per_second",
	Help:      "Count Visitors",
})

func main() {
	r := prometheus.NewRegistry()
	r.MustRegister(visitor)

	go func() {
		for {
			time.Sleep(time.Second)
			visitor.Set(0)
		}
	}()

	http.HandleFunc("/inc", func(w http.ResponseWriter, r *http.Request) {
		visitor.Inc()
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Done")
	})
	http.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	http.ListenAndServe(":8000", nil)
}
