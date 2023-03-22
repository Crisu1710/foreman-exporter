package metrics

import (
	"github.com/Crisu1710/foreman-exporter/pkg/collector"
	"github.com/Crisu1710/foreman-exporter/pkg/parser"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"time"
)

var reportTime = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "foreman_puppet_last_report",
	Help: "Timestamp of the last puppet run of each host",
}, []string{"host_name", "host_group_name", "ip", "status", "puppet_proxy_name"})

var interval = 1

func recordMetrics() {
	hosts := collector.GetHosts()
	reportTime.Reset()
	go func() {
		for _, host := range hosts.Results {
			if host.LastReport != "" {
				lastReport, err := parser.ConvertTime(host.LastReport)
				if err != nil {
					log.Println(err)
				}
				if host.HostGroupName == "" {
					host.HostGroupName = "None"
				}
				reportTime.WithLabelValues(host.Name, host.HostGroupName, host.Ip, host.GlobalStatusLabel, host.PuppetProxyName).Add(lastReport)
			}
		}
	}()
	interval = 15 * 60 // min to sec
}

func RunInterval() {
	for {
		time.Sleep(time.Duration(interval) * time.Second)
		recordMetrics()
	}
}
