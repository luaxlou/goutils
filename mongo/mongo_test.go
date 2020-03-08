package mongo

import "testing"

func TestNew(t *testing.T) {

	dsn := "mongodb://root:123456@localhost:27017"

	db := "testdb"

	n := New(dsn, db)

	n.Close()

}
