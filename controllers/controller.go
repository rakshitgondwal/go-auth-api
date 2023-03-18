package controllers

import (
	"github.com/labstack/echo/v4"
	"golang-auth/db"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUsers(client *mongo.Client ) echo.HandlerFunc {
	return func(c echo.Context) error {
		return db.FindAll("goapi-auth", "trial", client)
	}
}