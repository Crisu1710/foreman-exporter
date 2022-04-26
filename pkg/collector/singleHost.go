package collector

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type GetHost struct {
	GlobalStatusLabel string `json:"global_status_label"`
	LastReport        string `json:"last_report"`
	HostGroupName     string `json:"hostgroup_name"`
}

func SingleHost(id int) GetHost {
	strId := strconv.FormatInt(int64(id), 10)
	res, err := httpRequest("https://" + os.Getenv("FOREMAN_HOST") + "/api/hosts/" + strId)
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var host GetHost
	err = json.Unmarshal(body, &host)
	if err != nil {
		log.Fatal(err)
	}

	return host
}
