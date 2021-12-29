package influxdbv2

import influxdb2 "github.com/influxdata/influxdb-client-go/v2"

func New(url string, token string) *influxdb2.Client {

	c := influxdb2.NewClient(url, token)
	return &c
}
