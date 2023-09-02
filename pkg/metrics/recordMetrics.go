package metrics

import (
	"log"
	"time"

	"github.com/Crisu1710/foreman-exporter/pkg/collector"
	"github.com/Crisu1710/foreman-exporter/pkg/parser"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var reportTime = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "foreman_puppet_last_report",
		Help: "Timestamp of the last puppet run of each host",
	}, []string{"host_name", "host_group_name", "ip", "status", "puppet_proxy_name"})

func RecordMetrics() {
	hosts := collector.GetHosts()
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
				reportTime.WithLabelValues(host.Name, host.HostGroupName, host.Ip, host.GlobalStatusLabel, host.PuppetProxyName).Set(lastReport)
			}
			time.Sleep(15 * time.Second)
		}
	}()
}
