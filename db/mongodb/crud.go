package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneById(collectionName string, id primitive.ObjectID, obj interface{}, opts ...*options.FindOneOptions) error {

	filter := bson.D{
		{"_id", id},
	}

	return FindOne(collectionName, filter, obj, opts...)

}

func FindOne(collectionName string, filter interface{}, obj interface{}, opts ...*options.FindOneOptions) error {

	res := GetDB().Collection(collectionName).FindOne(context.Background(), filter, opts...)

	return res.Decode(obj)

}

func Find(collectionName string, filter interface{}, objs interface{}, opts ...*options.FindOptions) error {

	cur, err := GetDB().Collection(collectionName).Find(context.Background(), filter, opts...)

	if err != nil {
		return err
	}

	return cur.All(context.Background(), objs)

}

func UpdateOne(collectionName string, filter interface{}, updates interface{}, opts ...*options.UpdateOptions) error {

	_, err := GetDB().Collection(collectionName).UpdateOne(context.Background(), filter, updates, opts...)

	return err

}

func DeleteOne(collectionName string, filter interface{}, opts ...*options.DeleteOptions) error {

	_, err := GetDB().Collection(collectionName).DeleteOne(context.Background(), filter, opts...)

	return err

}

func DeleteMany(collectionName string, filter interface{}, opts ...*options.DeleteOptions) error {

	_, err := GetDB().Collection(collectionName).DeleteMany(context.Background(), filter, opts...)

	return err

}

func InsertOne(collectionName string, obj interface{}, opts ...*options.InsertOneOptions) (primitive.ObjectID, error) {

	res, err := GetDB().Collection(collectionName).InsertOne(context.Background(), obj, opts...)

	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), err

}

func Count(collectionName string, obj interface{}, opts ...*options.CountOptions) (int64, error) {

	n, err := GetDB().Collection(collectionName).CountDocuments(context.Background(), obj, opts...)

	if err != nil {
		return 0, err
	}

	return n, nil

}

func Aggregate(collectionName string, pipeline interface{}, objs interface{}, opts ...*options.AggregateOptions) error {

	cur, err := GetDB().Collection(collectionName).Aggregate(context.Background(), pipeline, opts...)

	if err != nil {
		return err
	}

	return cur.All(context.Background(), objs)

}

type SumResult struct {
	Total int64 `bson:"total"`
}

func Sum(collectionName string, columnName string, filter interface{}) int64 {

	pipeline := []bson.M{
		{"$match": filter},
		{"$group": bson.M{
			"_id":   nil,
			"total": bson.M{"$sum": "$" + columnName},
		}},
	}

	var res []SumResult

	err := Aggregate(collectionName, pipeline, &res)

	if err != nil || len(res) == 0 {
		return 0
	}

	return res[0].Total

}
