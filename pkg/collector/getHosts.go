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
	PuppetProxyName   string `json:"puppet_proxy_name"`
}

func GetHosts() GetAllHosts {
	res, err := httpRequest("https://" + os.Getenv("FOREMAN_HOST") + "/api/v2/hosts?per_page=1200")
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
