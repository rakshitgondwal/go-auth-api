package routes

import (
	"github.com/labstack/echo/v4"
	"../controllers/auth_controller.go"
)

func AuthRoutes(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/signup", controllers.Signup)
	auth.POST("/login", controllers.Login)
	auth.POST("/logout", controllers.Logout)
}
