package mysqldb

import "testing"

func TestNew(t *testing.T) {

	dsn := "root:123456@tcp(127.0.0.1:3306)/em?charset=utf8&parseTime=True&loc=Local&multiStatements=true"
	New(dsn)

}
