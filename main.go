package main

import (
	"log"
	"net/http"

	"github.com/Crisu1710/foreman-exporter/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	go metrics.RecordMetrics()
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		log.Println(err)
	}
}
