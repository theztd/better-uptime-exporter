package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	var (
		influx              bool
		prometheus, envPath string
	)

	flag.StringVar(&prometheus, "prometheus", "", "Path to the prometheus export file.")
	flag.BoolVar(&influx, "influx", false, "Send output to influx configured via ENV.")
	flag.StringVar(&envPath, "cfg", ".env", "Path to the .env style file.")
	flag.Parse()

	/*
		GET CONFIG
	*/
	godotenv.Load(envPath)
	BUP_TOKEN = os.Getenv("BUP_TOKEN")
	INFLUX_URL = os.Getenv("INFLUX_URL")
	INFLUX_TOKEN = os.Getenv("INFLUX_TOKEN")
	INFLUX_ORG = os.Getenv("INFLUX_ORG")
	INFLUX_BUCKET = os.Getenv("INFLUX_BUCKET")

	log.Println("INFO: Started to gathet url")

	data := listMetrics()

	if prometheus != "" {
		for _, m := range data {
			// better_uptime_metrics{url="", domain="Domain from url", monitor_type="", verify_ssl="true/false", method=""} STATUS
			fmt.Printf("better_uptime_metrics{url=\"%s\", domain=\"%s\", monitor_type=\"%s\", verify_ssl=\"%t\", method=\"%s\"} %d\n",
				m.Attr.Url, getDomainFromUrl(m.Attr.Url), m.Attr.Type, m.Attr.SSL, m.Attr.Method, printStatus(m.Attr.Status))
		}
	}

	if influx {
		for _, m := range data {

			metric := InfluxMetric{
				Name: "better_uptime_metrics",
				Tags: []InfluxTags{
					{Key: "Url", Value: m.Attr.Url},
					{Key: "Domain", Value: getDomainFromUrl(m.Attr.Url)},
					{Key: "Type", Value: m.Attr.Type},
					{Key: "SSL", Value: fmt.Sprintf("%t", m.Attr.SSL)},
					{Key: "Method", Value: m.Attr.Method},
					{Key: "StatusCode", Value: m.Attr.Status},
				},
				Fields: []InfluxFields{
					{Key: "Status", Value: float32(printStatus(m.Attr.Status))},
				},
			}

			err := influxWrite(metric)
			if err != nil {
				log.Println("ERR: [InfluxWrite]", m.Attr.Url, err)
			}

		}
	}

}
