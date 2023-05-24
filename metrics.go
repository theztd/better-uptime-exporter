package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type attributes struct {
	Name   string `json:"pronounceable_name"`
	Url    string `json:"url"`
	Type   string `json:"monitor_type"`
	SSL    bool   `json:"verify_ssl"`
	Method string `json:"http_method"`
	Status string `json:"status"`
}

type resultMetric struct {
	Type string     `json:"type"`
	Attr attributes `json:"attributes"`
}

func printStatus(input string) int {
	translate := map[string]int{
		"up":      0,
		"down":    1,
		"unknown": 9,
	}
	return translate[input]
}

func getDomainFromUrl(input string) string {
	x, _ := url.Parse(input)
	return x.Hostname()
}

func listMetrics() []resultMetric {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, "https://uptime.betterstack.com/api/v2/monitors", nil)
	req.Header.Add("Authorization", "Bearer "+BUP_TOKEN)
	res, err := client.Do(req)

	// Handle request errors
	if err != nil {
		log.Fatalln("ERR: [listMetrics]", err)
	}
	if res.Body == nil {
		log.Fatalln("ERR: [listMetrics] Empty result body")
	} else {
		defer res.Body.Close()
	}

	body, _ := ioutil.ReadAll(res.Body)
	var data = struct {
		Data []resultMetric `json:"data"`
		// Pagination `json:"pagination"`
	}{}
	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatalln("ERR: [listMetrics.ParseJson]", err)
	}

	log.Println(res.StatusCode)

	return data.Data
}
