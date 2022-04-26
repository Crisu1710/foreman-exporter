package collector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type GetData struct {
	Results []GetResults
}

type GetResults struct {
	Id   int
	Name string
}

func AllHosts() GetData {
	res, err := httpRequest("https://" + os.Getenv("FOREMAN_HOST") + "/api/hosts?thin=true&&per_page=1000")
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var data GetData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
