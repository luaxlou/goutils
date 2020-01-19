package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOneById(collectionName string, id primitive.ObjectID, obj interface{}, opts ...*options.FindOneOptions) error {

	filter := bson.D{
		{"_id", id},
	}

	log.Println(id)

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

func InsertOne(collectionName string, obj interface{}, opts ...*options.InsertOneOptions) (string, error) {

	res, err := GetDB().Collection(collectionName).InsertOne(context.Background(), obj, opts...)

	if err != nil {
		return "", err
	}

	return res.InsertedID.(primitive.ObjectID).Hex(), err

}
