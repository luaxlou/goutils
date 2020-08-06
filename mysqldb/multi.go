package mysqldb

import (
	"github.com/jinzhu/gorm"
)

//使用多例模式======

var dsns = make(map[string]string, 0)
var instances = make(map[string]*gorm.DB, 0)

func Add(name, dsn string) {

	dsns[name] = dsn

}

func GetInstance(name string) *gorm.DB {

	ins, ok := instances[name]

	if ok {
		return ins
	}

	dsn, ok1 := dsns[name]

	if !ok1 {

		return nil
	}

	ins = New(dsn)

	instances[name] = ins

	return ins
}
