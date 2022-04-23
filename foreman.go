package main

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type getHost struct {
	StatusLabel string `json:"status_label"`
	Status      float64
}

type getData struct {
	Total float64
	Page  int
	//Results map[int][]GetResults
	Results []GetResults
}

type GetResults struct {
	Id   int
	Name string
}

func httpRequest(url string) (*http.Response, error) {
	metaClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := metaClient.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func allHosts() getData {
	res, err := httpRequest("http://localhost:8081/api/hosts.json")

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var data getData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func singleHost(id int) getHost {
	newId := strconv.FormatInt(int64(id), 10)
	res, err := httpRequest("http://localhost:8081/api/hosts/" + newId + "/status/global.json")

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var host getHost
	err = json.Unmarshal(body, &host)
	if err != nil {
		log.Fatal(err)
	}

	return host
}

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

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
