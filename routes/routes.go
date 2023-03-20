package routes

import (
	"golang-auth/controllers"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, client *mongo.Client) {
	// //Public Routes

	// //Admin Routes
	// e.DELETE("/delete/:id", DeleteUser)

	// //User Routes
	// e.GET("/", controllers.GetUsers)
	e.GET("/", controllers.GetUsers(client))
	e.POST("/register", controllers.CreateUser(client))

	e.POST("/login", controllers.LoginUser(client))
	e.POST("/logout", controllers.LogoutUser(client))
	e.POST("/refresh", controllers.RefreshToken(client))

}
