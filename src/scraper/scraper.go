package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

func scrape(url string, timeout int) ([]string, error) {
	proxyRegex := regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?):([0-9]){1,4}`)
	client := &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
	response, err := client.Get(url)
	if err != nil {
		return nil, errors.New("bad request")
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("could not read body")
	}
	proxies := proxyRegex.FindAllString(string(body), -1)
	return proxies, nil
}
