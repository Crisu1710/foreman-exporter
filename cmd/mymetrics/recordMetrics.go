package mymetrics

import (
	"example.com/mycollector"
	"example.com/mytime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"time"
)

var forTarget = promauto.NewGaugeVec(prometheus.GaugeOpts{
	Name: "foreman_puppet_last_check",
	Help: "Timestamp of the last puppet run of each host",
}, []string{"host_name", "status"})

// DEV []string{"host_id", "host_name", "status", "status_id"}

var interval = 1

func recordMetrics() {
	data := mycollector.AllHosts()
	go func() {
		for _, d := range data.Results {
			host := mycollector.SingleHost(d.Id)
			if host.LastReport != "" {
				newtime := mytime.ConvertTime(host.LastReport)
				// DEV hostId := strconv.Itoa(d.Id)
				// DEV statusId := int64(host.GlobalStatus)
				// DEV statusIdStr := strconv.FormatInt(statusId, 10)
				forTarget.WithLabelValues(d.Name, host.GlobalStatusLabel).Add(newtime)
				// DEV forTarget.WithLabelValues(hostId, d.Name, host.GlobalStatusLabel, statusIdStr).Add(newtime)
			}
		}
	}()
	interval = 15
}

func RunInterval() {
	for {
		time.Sleep(time.Duration(interval) * time.Minute)
		go recordMetrics()
	}
}
