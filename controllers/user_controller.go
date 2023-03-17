package controllers

import(
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"../routes/user_route.go"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetUsers(ctx echo.Context) error {
	users, err := services.GetAllUsers()
	if err!= nil{
		return ctx.JSON(http.StatusInternalServerError, "Failed to get users")
	}
	return ctx.JSON(http.StatusOK, users)
}