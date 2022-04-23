package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func recordMetrics() {
	data := allHosts()
	go func() {
		totalTargets.Add(data.Total)
		for _, d := range data.Results {
			host := singleHost(d.Id)
			forTarget.WithLabelValues(d.Name, host.StatusLabel).Add(host.Status)
		}
	}()
}

var totalTargets = promauto.NewCounter(prometheus.CounterOpts{
	Name: "foreman_total_targets",
	Help: "test bla",
})

var forTarget = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "foreman_puppet_host",
	Help: "test bla",
}, []string{"host", "status"})
