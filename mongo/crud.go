package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *DB) FindOneById(collectionName string, id primitive.ObjectID, obj interface{}, opts ...*options.FindOneOptions) error {

	filter := bson.D{
		{"_id", id},
	}

	return c.FindOne(collectionName, filter, obj, opts...)

}

func (c *DB) FindOne(collectionName string, filter interface{}, obj interface{}, opts ...*options.FindOneOptions) error {

	res := c.DB.Collection(collectionName).FindOne(context.Background(), filter, opts...)

	return res.Decode(obj)

}

func (c *DB) Find(collectionName string, filter interface{}, objs interface{}, opts ...*options.FindOptions) error {

	cur, err := c.DB.Collection(collectionName).Find(context.Background(), filter, opts...)

	if err != nil {
		return err
	}

	return cur.All(context.Background(), objs)

}

func (c *DB) UpdateOne(collectionName string, filter interface{}, updates interface{}, opts ...*options.UpdateOptions) error {

	_, err := c.DB.Collection(collectionName).UpdateOne(context.Background(), filter, updates, opts...)

	return err

}

func (c *DB) DeleteOne(collectionName string, filter interface{}, opts ...*options.DeleteOptions) error {

	_, err := c.DB.Collection(collectionName).DeleteOne(context.Background(), filter, opts...)

	return err

}

func (c *DB) DeleteMany(collectionName string, filter interface{}, opts ...*options.DeleteOptions) error {

	_, err := c.DB.Collection(collectionName).DeleteMany(context.Background(), filter, opts...)

	return err

}

func (c *DB) InsertOne(collectionName string, obj interface{}, opts ...*options.InsertOneOptions) (primitive.ObjectID, error) {

	res, err := c.DB.Collection(collectionName).InsertOne(context.Background(), obj, opts...)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), err

}

func (c *DB) Count(collectionName string, obj interface{}, opts ...*options.CountOptions) (int64, error) {

	n, err := c.DB.Collection(collectionName).CountDocuments(context.Background(), obj, opts...)

	if err != nil {
		return 0, err
	}

	return n, nil

}

func (c *DB) Aggregate(collectionName string, pipeline interface{}, objs interface{}, opts ...*options.AggregateOptions) error {

	cur, err := c.DB.Collection(collectionName).Aggregate(context.Background(), pipeline, opts...)

	if err != nil {
		return err
	}

	return cur.All(context.Background(), objs)

}

type SumResult struct {
	Total int64 `bson:"total"`
}

func (c *DB) Sum(collectionName string, columnName string, filter interface{}) int64 {

	pipeline := []bson.M{
		{"$match": filter},
		{"$group": bson.M{
			"_id":   nil,
			"total": bson.M{"$sum": "$" + columnName},
		}},
	}

	var res []SumResult

	err := c.Aggregate(collectionName, pipeline, &res)

	if err != nil || len(res) == 0 {
		return 0
	}

	return res[0].Total

}
