package metrics

import (
	"example.com/collector"
	"example.com/parser"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var forTarget = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "foreman_puppet_last_check",
	Help: "Timestamp of the last puppet run of each host",
}, []string{"host_name", "host_group_name", "status"})

var interval = 1

func recordMetrics() {
	data := collector.AllInOneHosts()
	forTarget.Reset()
	go func() {
		for _, d := range data.Results {
			if d.LastReport != "" {
				newtime := parser.ConvertTime(d.LastReport)
				if d.HostGroupName == "" {
					d.HostGroupName = "None"
				}
				forTarget.WithLabelValues(d.Name, d.HostGroupName, d.GlobalStatusLabel).Add(newtime)
			}
		}
	}()
	interval = 20 // * 60 // min to sec
}

func RunInterval() {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		recordMetrics()
	}
}
