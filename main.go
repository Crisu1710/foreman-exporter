package main

import (
	"example.com/mymetrics"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	go mymetrics.RunInterval()

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(":2112", nil)
	if err != nil {
		fmt.Println(err)
	}
}
