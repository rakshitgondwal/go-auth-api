package mdlware

import (
	"fmt"
	"golang-auth/configs"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	// "github.com/your-username/your-app/models"
)

// RequireLogin middleware function to require a valid JWT token for protected routes
func RequireLogin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Verify the token signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            // Load the secret key from a secure location
            secretKey := []byte(configs.Cfg.JwtSecret)

            return secretKey, nil
        })
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        // Check if the token is valid
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
        c.Set("userID", claims["userID"])

        return next(c)
    }
}

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Get the JWT token from the request header
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return echo.NewHTTPError(http.StatusUnauthorized, "Missing token")
        }

        // Parse and validate the JWT token
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Verify the token signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            // Load the secret key from a secure location
            secretKey := []byte(configs.Cfg.JwtSecret)

            return secretKey, nil
        })
        if err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        // Check if the token is valid
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        // Check if the user has the "admin" role in their JWT claims
        if role, ok := claims["role"].(string); !ok || role != "admin" {
            return echo.NewHTTPError(http.StatusForbidden, "You are not authorized to access this resource")
        }

        // Set the user ID from the token claims in the context for later use
        c.Set("userID", claims["userID"])

        return next(c)
    }
}
