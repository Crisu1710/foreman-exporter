package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)

type getHost struct {
	StatusLabel string `json:"status_label"`
	Status      float64
}

func singleHost(id int) getHost {
	newId := strconv.FormatInt(int64(id), 10)
	res, err := httpRequest("https://foreman.example.com/api/hosts/" + newId + "/status/global")
	if err != nil {
		log.Fatal(err)
	}

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
