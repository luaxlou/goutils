package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	Client *mongo.Client

	DB *mongo.Database
}

func New(dsn, db string) *DB {

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		panic(err.Error())
	}

	mdb := client.Database(db)

	return &DB{
		Client: client,
		DB:     mdb,
	}
}

func (c *DB) Close() {
	c.Client.Disconnect(context.Background())

}
