package routes

import (
	"github.com/labstack/echo/v4"
	"golang-auth/controllers"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, client *mongo.Client){
	// //Public Routes
	// e.POST("/login", LoginUser)
	// e.POST("/logout", LogoutUser)
	// e.POST("/refresh", RefreshToken)

	// //Admin Routes
	// e.DELETE("/delete/:id", DeleteUser)

	// //User Routes
	// e.GET("/", controllers.GetUsers)
	e.GET("/", controllers.GetUsers(client))
	e.POST("/register", controllers.CreateUser(client))
}