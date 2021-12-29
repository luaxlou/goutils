package influxdbv2

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

func New(url string, token string) influxdb2.Client {

	return influxdb2.NewClient(url, token)
}

func Close() {

	instances.Range(func(key, value interface{}) bool {

		g := value.(influxdb2.Client)

		if g != nil {
			g.Close()
		}

		return true

	})
}
