package main

import (
	"net/http"
	"time"
)

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
