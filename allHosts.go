package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

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
