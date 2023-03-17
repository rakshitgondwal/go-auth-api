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

//GetUsers retrives all the users.

func (c *UserController) GetUsers(ctx echo.Context) error {
	users, err := services.GetAllUsers()
	if err!= nil{
		return ctx.JSON(http.StatusInternalServerError, "Failed to get users")
	}
	return ctx.JSON(http.StatusOK, users)
}

//GetUserByID retrieves a specific user by ID.

func (c *UserController) GetUserByID(ctx echo.Context) error{
	userID, err := strconv.Atoi(ctx.Param("id"))

	if err != nil{
		return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	user, err := services.GetUserByID(userID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "User not found")
	}

	return ctx.JSON(http.StatusOK, user)
}

//CreateUser creates a new user.

func(c *UserController) CreateUser(ctx echo.Context) error {
	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid user data")
	}

	if err := services.CreateUser(&user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to create user")
	}

	return ctx.JSON(http.StatusCreated, user)
}

//UpdateUser updates an existing user.
func(c *UserController) UpdateUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := ctx.Bind(&user); err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid user data")
	}

	if err := services.UpdateUser(userID, &user); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to update user")
	}

	return ctx.NoContent(http.StatusOK)
}

// DeleteUser deletes an existing user.
func (c *UserController) DeleteUser(ctx echo.Context) error {
	userID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	if err := services.DeleteUser(userID); err != nil {
		return ctx.JSON(http.StatusInternalServerError, "Failed to delete user")
	}

	return ctx.NoContent(http.StatusOK)
}