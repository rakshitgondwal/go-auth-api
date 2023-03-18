package db

import (
	"context"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll(db string, collection string,client *mongo.Client) (error) {
	coll := client.Database(db).Collection(collection)
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return err
	}
	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		return err
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}
	return nil
}