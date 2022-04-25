package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"time"
)

var forTarget = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "foreman_puppet_last_check",
	Help: "Timestamp of the last puppet run of each host",
}, []string{"host_id", "host_name", "status", "status_id"})

var testInterval = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "foreman_test_interval",
	Help: "test",
}, []string{"test"})

var interval = 1

func recordMetrics() {
	data := allHosts()
	go func() {
		testInterval.WithLabelValues("bla").Inc()
		for _, d := range data.Results {
			host := singleHost(d.Id)
			if host.LastReport != "" {
				newtime := convertTime(host.LastReport)
				hostId := strconv.Itoa(d.Id)
				statusId := int64(host.GlobalStatus)
				statusIdStr := strconv.FormatInt(statusId, 10)
				forTarget.WithLabelValues(hostId, d.Name, host.GlobalStatusLabel, statusIdStr).Add(newtime)
			}
		}
	}()
	interval = 15
}

func runInterval() {
	for {
		time.Sleep(time.Duration(interval) * time.Minute)
		go recordMetrics()
	}
}
