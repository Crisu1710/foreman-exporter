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

var interval = 5

func recordMetrics() {
	data := collector.AllHosts()
	go func() {
		for _, d := range data.Results {
			host := collector.SingleHost(d.Id)
			if host.LastReport != "" {
				newtime := parser.ConvertTime(host.LastReport)
				forTarget.WithLabelValues(d.Name, host.HostGroupName, host.GlobalStatusLabel).Add(newtime)
			}
		}
	}()
	interval = 15 * 60 // min to sec
}

func RunInterval() {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		go recordMetrics()
	}
}
