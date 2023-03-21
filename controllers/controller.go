package controllers

import (
	// "context"
	"fmt"
	"golang-auth/db"
	"net/http"
	"strconv"
	"time"

	"golang-auth/jwt"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the request body to get the user credentials
        username := c.QueryParam("username")
		password := c.QueryParam("password")

		// Check if the user exists in the database
		dbUser, err := db.FindOne(username, "goapi-auth", "users", client)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if dbUser == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// Check if the provided password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// Generate a JWT token for the authenticated user
		token, err := jwt.GenerateToken(dbUser.Username)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		loc, err := time.LoadLocation("Asia/Kolkata")
		if err != nil {
			fmt.Println(err)
			return err
		}
		
		newToken := db.Tokens{
			Token: token,
			Username: username,
			CreatedAt: time.Now().In(loc),
			UpdatedAt: time.Now().In(loc),
			ExpiresAt: time.Now().In(loc).Add(1 * time.Hour),
		}

		db.AddToken(newToken, client)
		fmt.Print(newToken)

		// Store the JWT token in the response header
		c.Response().Header().Set("Authorization", "Bearer "+token)

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func LogoutUser(client *mongo.Client) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Get the token from the Authorization header
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return echo.NewHTTPError(http.StatusBadRequest, "Missing token in request header")
        }

        // Revoke the token by adding it to the blacklist
        err := jwt.RevokeToken(tokenString, client)
        if err != nil {
            return echo.NewHTTPError(http.StatusInternalServerError, "Failed to revoke token")
        }

        return c.JSON(http.StatusOK, "Successfully logged out")
    }
}

func RefreshToken(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the request body to get the user credentials
		username := c.QueryParam("username")
		password := c.QueryParam("password")

		// Check if the user exists in the database
		dbUser, err := db.FindOne(username, "goapi-auth", "users", client)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		if dbUser == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// Check if the provided password is correct
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(password)); err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid username or password")
		}

		// Generate a JWT token for the authenticated user
		token, err := jwt.GenerateToken(dbUser.Username)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		err2 := db.UpdateToken(dbUser.Username,client,token,c)
		if err2 != nil{
			return echo.NewHTTPError(http.StatusInternalServerError, err2.Error())
		}
		fmt.Println("new token:")
		fmt.Println(token)

		// Store the JWT token in the response header
		c.Response().Header().Set("Authorization", "Bearer "+token)

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

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

func CreateUser(client *mongo.Client) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse the incoming data from Postman
		ID := primitive.NewObjectID()
		username := c.QueryParam("username")
		password := c.QueryParam("password")
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		isAdmin, err := strconv.ParseBool(c.QueryParam("isAdmin"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid value for 'isAdmin'"})
		}
		organization := c.QueryParam("organization")
		user := db.User{ID: ID, Username: username, Password: string(hashedPassword), IsAdmin: isAdmin, Organization: organization}
		// Insert the data into the database
		return db.AddUser(user, client, c)
	}
}
