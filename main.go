package main

import (
	"example.com/metrics"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	go metrics.RunInterval()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		fmt.Println(err)
	}
}
