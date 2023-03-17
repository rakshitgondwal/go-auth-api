package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// LoginUser is used for user authentication
func LoginUser(c echo.Context) error {
	// Get the login credentials from the request body
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Call the authentication service to validate the user credentials
	if AuthenticateUser(username, password) {
		// Generate and return a JWT token
		token := GenerateToken(username)
		return c.JSON(http.StatusOK, map[string]string{
			"message": "User logged in successfully",
			"token":   token,
		})
	}

	// Return an error if the authentication failed
	return echo.ErrUnauthorized
}

// AuthenticateUser checks if the provided user credentials are valid or not
func AuthenticateUser(username, password string) bool {
	// Call your authentication logic here to check if the user credentials are valid
	// For example, you can check if the username and password exist in the database or not
	// If the user credentials are valid, return true, otherwise, return false
	// This is just a dummy implementation for demonstration purposes
	if username == "admin" && password == "password" {
		return true
	}
	return false
}

// GenerateToken generates a JWT token for the authenticated user
func GenerateToken(username string) string {
	// Call your JWT token generation logic here
	// For example, you can use the "github.com/dgrijalva/jwt-go" package to generate a JWT token
	// This is just a dummy implementation for demonstration purposes
	return "dummy_token_for_" + username
}