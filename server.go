package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

func httpRequest(url string) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	metaClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Cookie", "sid=XXXXXXXXXX; _session_id=XXXXXXXXXX; timezone=Europe%2FBerlin")

	res, err := metaClient.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
