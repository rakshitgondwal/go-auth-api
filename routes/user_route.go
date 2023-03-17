package routes

import (
	"github.com/labstack/echo/v4"
	"../controllers/user_controller.go"
	"../middleware/auth_controller.go"
)

func UserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	// Public routes
	userGroup.POST("/register", controllers.RegisterUser)

	// Auth routes
	authGroup := userGroup.Group("")
	authGroup.Use(middleware.JWTMiddleware)
	authGroup.GET("", controllers.ListUsers)
	authGroup.GET("/:id", controllers.GetUser)
	authGroup.PUT("/:id", controllers.UpdateUser)
	authGroup.DELETE("/:id", controllers.DeleteUser)
}