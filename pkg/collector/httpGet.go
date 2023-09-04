package collector

import (
	"crypto/tls"
	"net/http"
	"os"
	"time"
)

func httpRequest(url string) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	metaClient := http.Client{
		Timeout: 20 * time.Second,
	}

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Authorization", os.Getenv("FOREMAN_PW"))

	res, err := metaClient.Do(request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
