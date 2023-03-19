package controllers

import (
	"golang-auth/db"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(client *mongo.Client) echo.HandlerFunc {
    return func(c echo.Context) error {
        results, err := db.FindAll("goapi-auth", "trial", client)
        if err != nil {
            return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
        }
        return c.JSON(http.StatusOK, results)
    }
}