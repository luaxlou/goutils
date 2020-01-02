package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"strings"
)

var mgo *mongo.Client

var mdb *mongo.Database

func Init() {
	dsn := os.Getenv("MONGO_DSN")
	db := os.Getenv("MONGO_DB")

	if dsn == "" {
		panic("ENV value MONGO_DSN required. example: 'mongodb://root:123456@localhost:27017'")

	}

	if db == "" {
		panic("ENV value MONGO_DB required.")

	}

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err.Error())
	}

	mdb = client.Database(db)
	mgo = client
}

func GetDB() *mongo.Database {

	if mgo == nil {

		//Lazy init
		Init()

	}

	return mdb

}

func Close() {
	mgo.Disconnect(context.Background())

}

func IsDuplicateError(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "dup") {
			return true
		}
	}

	return false
}
