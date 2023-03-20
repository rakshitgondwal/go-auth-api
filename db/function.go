package db

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
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

func AddUser(user User, client *mongo.Client, c echo.Context) error{
    coll := client.Database("goapi-auth").Collection("users")
        _, err := coll.InsertOne(context.Background(), user)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        }

        return c.JSON(http.StatusOK, map[string]string{"message": "User added successfully"})
}

func AddToken(token Tokens, client *mongo.Client) error {
    coll := client.Database("goapi-auth").Collection("tokens")
    _, err := coll.InsertOne(context.Background(), token)
    return err
}

func AddTokenToBlacklist(token RevokedToken, client *mongo.Client) error {
    coll := client.Database("goapi-auth").Collection("blacklisted-tokens")
    _, err := coll.InsertOne(context.Background(), token)
    return err
}

func FindOne(username string, db string, collection string, client *mongo.Client) (*User, error) {
    var user User
    coll := client.Database(db).Collection(collection)
    err := coll.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, err
    }
    return &user, nil
}