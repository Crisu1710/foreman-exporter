package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"strconv"
	"time"
)

func convertTime(newtime string) float64 {
	layout := "2006-01-02 15:04:05 UTC"
	t, err := time.Parse(layout, newtime)
	if err != nil {
		fmt.Println(err)
	}
	tofloat64 := float64(t.Unix())

	return tofloat64
}

func recordMetrics() {
	data := allHosts()
	go func() {
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
}

var forTarget = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "foreman_puppet_last_check",
	Help: "Timestamp of the last puppet run of each host",
}, []string{"host_id", "host_name", "status", "status_id"})
