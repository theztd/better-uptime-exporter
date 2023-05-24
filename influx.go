package main

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxTags struct {
	Key   string
	Value string
}

type InfluxFields struct {
	Key   string
	Value float32
}

type InfluxMetric struct {
	Name   string
	Tags   []InfluxTags
	Fields []InfluxFields
}

func influxWrite(metricData InfluxMetric) error {
	client := influxdb2.NewClient(INFLUX_URL, INFLUX_TOKEN)
	writeApi := client.WriteAPIBlocking(INFLUX_ORG, INFLUX_BUCKET)

	p := influxdb2.NewPointWithMeasurement(metricData.Name)

	for _, tag := range metricData.Tags {
		p.AddTag(tag.Key, tag.Value)
	}

	for _, item := range metricData.Fields {
		p.AddField(item.Key, item.Value)
	}

	p.SetTime(time.Now())

	err := writeApi.WritePoint(context.Background(), p)
	if err != nil {
		return err
	}

	return nil
}
