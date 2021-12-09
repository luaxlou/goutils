package redisdb

import (
	"github.com/go-redis/redis/v8"
	"sync"
)

//使用多例模式======

var instances sync.Map

func Add(name, host string, pass string, db int) {

	_, ok := instances.Load(name)

	if ok {
		return
	}

	instances.Store(name, New(host, pass, db))

}

func GetInstance(name string) *redis.Client {

	ins, ok := instances.Load(name)

	if ok {
		return ins.(*redis.Client)
	}

	return nil
}
