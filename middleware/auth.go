package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/your-username/your-app/models"
)

// RequireLogin middleware function to require a valid JWT token for protected routes
func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get JWT token from request header
		tokenString := c.Request().Header.Get("Authorization")
		if tokenString == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing or invalid JWT token")
		}

		// Verify JWT token
		token, err := models.VerifyToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid JWT token")
		}

		// Attach token claims to request context for further processing
		c.Set("token", token)

		return next(c)
	}
}

// RequireAdmin middleware function to require the authenticated user to have admin permissions
func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get user ID from token claims
		token := c.Get("token").(*models.Token)
		userID := token.UserID

		// Check if user has admin permissions
		user, err := models.GetUserByID(userID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to retrieve user information")
		}
		if !user.IsAdmin {
			return echo.NewHTTPError(http.StatusForbidden, "User does not have admin permissions")
		}

		return next(c)
	}
}
