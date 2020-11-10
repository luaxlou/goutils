package mysqldb

import (
	"sync"

	"gorm.io/gorm"
)

//使用多例模式======

var dsns = make(map[string]string, 0)
var instances sync.Map

func Add(name, dsn string) {

	dsns[name] = dsn

}

func GetInstance(name string) *gorm.DB {

	ins, ok := instances.Load(name)

	if ok {
		return ins.(*gorm.DB)
	}

	dsn, ok1 := dsns[name]

	if !ok1 {

		return nil
	}

	ins = New(dsn)

	instances.Store(name, ins)

	return ins.(*gorm.DB)
}
