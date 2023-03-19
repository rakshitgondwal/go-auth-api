package main

import (
	"context"
	"fmt"

	// "net/http"

	"golang-auth/configs"
	"golang-auth/db"
	"golang-auth/routes"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

var e = echo.New()
var client *mongo.Client

func init() {
	//Initialize the cleanenv package
	err := cleanenv.ReadEnv(&configs.Cfg)
	fmt.Printf("%v", configs.Cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
	
	//Setup database connection
	cl, err := db.ConnectDB()
	if err != nil {
		e.Logger.Fatal("Unable to connect to database")
	}
	
	//Set up the routes
	routes.InitRoutes(e, cl)
	client = cl
}

func main() {
	
	coll := client.Database("goapi-auth").Collection("users")
	users := []interface{}{
		db.User{Username: "rakshitgondwal", Password: "rakshitgondwal", IsAdmin: false, Organization: "first"},
		db.User{Username: "notrakshit", Password: "notrakshit", IsAdmin: true, Organization: "second"},
	}
	_, err2 := coll.InsertMany(context.TODO(), users)
	if err2 != nil {
		e.Logger.Fatal("Unable to add data to the database")
	}

	//Start the server
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", configs.Cfg.Port)))
}
