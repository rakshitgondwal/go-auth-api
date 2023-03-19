package controllers

import (
	// "context"
	"golang-auth/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// func generateJWT(){

// }

func GetUsers(client *mongo.Client) echo.HandlerFunc {
    // 1. Is user logged in or not
    // 2. Check for org
    return func(c echo.Context) error {
        results, err := db.FindAll("goapi-auth", "users", client)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusOK, results)
    }
}

func CreateUser(client *mongo.Client) echo.HandlerFunc{
    return func(c echo.Context) error {
        // Parse the incoming data from Postman
        username := c.QueryParam("username")
        password := c.QueryParam("password")
        hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password), 8)
        if err != nil{
            return c.JSON(http.StatusBadRequest,err)
        }
        isAdmin, err := strconv.ParseBool(c.QueryParam("isAdmin"))
        if err != nil {
            return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for 'isAdmin'"})
        }
        organization := c.QueryParam("organization")
        user := db.User{Username: username, Password: string(hashedPassword), IsAdmin: isAdmin, Organization: organization}
        // Insert the data into the database
         return db.AddUser(user, client, c)
    }
}