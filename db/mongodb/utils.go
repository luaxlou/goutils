package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindOneById(collectionName string, idHex string, obj interface{}) error {

	id, _ := primitive.ObjectIDFromHex(idHex)

	filter := bson.D{
		{"_id", id},
	}

	log.Println(id)

	return FindOne(collectionName, filter, obj)

}

func FindOne(collectionName string, filter interface{}, obj interface{}) error {

	res := GetDB().Collection(collectionName).FindOne(context.Background(), filter)

	return res.Decode(obj)

}

func Find(collectionName string, filter interface{}, objs interface{}) error {

	cur, err := GetDB().Collection(collectionName).Find(context.Background(), filter)

	if err != nil {
		return err
	}

	return cur.All(context.Background(), objs)

}

func UpdateOne(collectionName string, filter interface{}, updates interface{}) error {

	_, err := GetDB().Collection(collectionName).UpdateOne(context.Background(), filter, updates)

	return err

}

func DeleteOne(collectionName string, filter interface{}) error {

	_, err := GetDB().Collection(collectionName).DeleteOne(context.Background(), filter)

	return err

}

func InsertOne(collectionName string, obj interface{}) (string, error) {

	res, err := GetDB().Collection(collectionName).InsertOne(context.Background(), obj)

	if err != nil {
		return "", err
	}

	return res.InsertedID.(string), err

}
