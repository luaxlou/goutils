package influxdbv2

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"sync"
)

//使用多例模式======

var instances sync.Map

func Add(name string, url string, token string) {

	_, ok := instances.Load(name)

	if ok {
		return
	}

	instances.Store(name, New(url, token))

}

func GetInstance(name string) *influxdb2.Client {

	ins, ok := instances.Load(name)

	if ok {
		return ins.(*influxdb2.Client)
	}

	return nil
}
