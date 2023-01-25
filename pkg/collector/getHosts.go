package collector

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type GetAllHosts struct {
	Results []GetAllResults
}

type GetAllResults struct {
	GlobalStatusLabel string `json:"global_status_label"`
	LastReport        string `json:"last_report"`
	HostGroupName     string `json:"hostgroup_name"`
	Name              string `json:"name"`
	Ip                string `json:"ip"`
}

func GetHosts() GetAllHosts {
	res, err := httpRequest("https://" + os.Getenv("FOREMAN_HOST") + "/api/hosts?per_page=1000")
	if err != nil {
		log.Println(err)
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr)
	}

	var data GetAllHosts
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Println(err)
	}

	return data
}
