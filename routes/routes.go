package routes

import (
	"golang-auth/controllers"
	"golang-auth/mdlware"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(e *echo.Echo, client *mongo.Client) {
	
	//PUBLIC ROUTES
	e.POST("/login", controllers.LoginUser(client))
	e.POST("/logout", controllers.LogoutUser(client))
	e.POST("/refresh", controllers.RefreshToken(client))
	
	//USER ROUTES
	e.GET("/", controllers.GetUsers(client), mdlware.RequireLogin)

	//ADMIN ROUTES
	e.POST("/register", controllers.CreateUser(client), mdlware.RequireAdmin)
	e.POST("/delete", controllers.DeleteUser(client), mdlware.RequireAdmin)

}
