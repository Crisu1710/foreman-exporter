package collector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type GetAllData struct {
	Results []GetAllResults
}

type GetAllResults struct {
	GlobalStatusLabel string `json:"global_status_label"`
	LastReport        string `json:"last_report"`
	HostGroupName     string `json:"hostgroup_name"`
	Name              string `json:"name"`
}

func AllInHosts() GetAllData {
	res, err := httpRequest("https://" + os.Getenv("FOREMAN_HOST") + "/api/hosts?per_page=1000")
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var data GetAllData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
