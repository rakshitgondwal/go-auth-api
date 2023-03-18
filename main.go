package main

import (
	"context"
	"encoding/json"
	"fmt"

	// "net/http"

	"golang-auth/configs"
	"golang-auth/db"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

var e = echo.New()

func init() {
	err := cleanenv.ReadEnv(&configs.Cfg)
	fmt.Printf("%v", configs.Cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

func main() {

	client, err := db.ConnectDB()
	if err != nil {
		e.Logger.Fatal("Unable to connect to database")
	}

	//MongoDB schema which is in db/models.go
	type Tea struct {
		Type     string
		Category string
		Toppings []string
		Price    float32
	}

	coll := client.Database("goapi-auth").Collection("trial")
	docs := []interface{}{
		Tea{Type: "Masala", Category: "black", Toppings: []string{"ginger", "pumpkin spice", "cinnamon"}, Price: 6.75},
		Tea{Type: "Gyokuro", Category: "green", Toppings: []string{"berries", "milk foam"}, Price: 5.65},
		Tea{Type: "English Breakfast", Category: "black", Toppings: []string{"whipped cream", "honey"}, Price: 5.75},
		Tea{Type: "Sencha", Category: "green", Toppings: []string{"lemon", "whipped cream"}, Price: 5.15},
		Tea{Type: "Assam", Category: "black", Toppings: []string{"milk foam", "honey", "berries"}, Price: 5.65},
		Tea{Type: "Matcha", Category: "green", Toppings: []string{"whipped cream", "honey"}, Price: 6.45},
		Tea{Type: "Earl Grey", Category: "black", Toppings: []string{"milk foam", "pumpkin spice"}, Price: 6.15},
		Tea{Type: "Hojicha", Category: "green", Toppings: []string{"lemon", "ginger", "milk foam"}, Price: 5.55},
	}
	_, err2 := coll.InsertMany(context.TODO(), docs)
	if err2 != nil {
		e.Logger.Fatal("Unable to add data to the database")
	}

	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		panic(err)
	}
	var results []Tea
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	for _, result := range results {
		res, _ := json.Marshal(result)
		fmt.Println(string(res))
	}

	// err := cleanenv.ReadEnv(&configs.Cfg)
	// if err != nil {
	// 	e.Logger.Fatal("Unable to load configuration")
	// }

	// e.POST("/save", func(c echo.Context) error {
	// 	name := c.FormValue("name")
	// 	email := c.FormValue("email")
	// 	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
	// })

	// e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", configs.Cfg.Port)))
}
