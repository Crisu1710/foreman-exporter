package main

import (
	"github.com/Crisu1710/foreman-exporter/pkg/metrics"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	go metrics.RunInterval()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		log.Println(err)
	}
}
