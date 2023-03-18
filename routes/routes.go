package routes

import (
	"github.com/labstack/echo/v4"
	"golang-auth/controllers"
)

func InitRoutes(e *echo.Echo){
	// //Public Routes
	// e.POST("/login", LoginUser)
	// e.POST("/logout", LogoutUser)
	// e.POST("/refresh", RefreshToken)

	// //Admin Routes
	// adminGroup := e.Group("/admins")
	// adminGroup.PUT("/register", CreateUser)
	// adminGroup.DELETE("/delete/:id", DeleteUser)

	// //User Routes
	// userGroup := e.Group("/users")
	// userGroup.GET("/", controllers.GetUsers)
	e.GET("/", controllers.GetUsers)
}