package routes

import (
	"golang-auth/controllers"
	"golang-auth/middleware"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, client *mongo.Client) {
	
	//PUBLIC ROUTES
	e.POST("/register", controllers.CreateUser(client))
	e.POST("/login", controllers.LoginUser(client))
	e.POST("/logout", controllers.LogoutUser(client))
	e.POST("/refresh", controllers.RefreshToken(client))
	
	//USER ROUTES
	e.GET("/", controllers.GetUsers(client), middleware.RequireLogin)

	//ADMIN ROUTES
	e.POST("/add", controllers.CreateUser(client))
	e.POST("/delete", controllers.DeleteUser(client))

}
