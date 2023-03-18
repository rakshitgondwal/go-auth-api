package db

import (
	"context"
    "fmt"
	"time"

    "golang-auth/configs"

    "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
    
)

func ConnectDB() (*mongo.Client, error) {

    serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
    uri := fmt.Sprintf("mongodb+srv://%s:%s@goapi-auth.nclmvto.mongodb.net/?retryWrites=true&w=majority", configs.Cfg.DbUsername, configs.Cfg.DbPassword)
    clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }
    var result bson.M
    if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
        return nil, err
    }
    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
    return client, nil
}

