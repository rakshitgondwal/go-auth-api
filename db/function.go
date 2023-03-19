package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindAll(db string, collection string, client *mongo.Client) ([]User, error) {
    coll := client.Database(db).Collection(collection)
    cursor, err := coll.Find(context.TODO(), bson.D{})
    if err != nil {
        return nil, err
    }
    var results []User
    if err = cursor.All(context.TODO(), &results); err != nil {
        return nil, err
    }
    return results, nil
}