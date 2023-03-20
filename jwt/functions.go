package jwt

import (
	"golang-auth/configs"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateToken(username string) (string, error) {
    // Set the expiration time for the token
    expirationTime := time.Now().Add(1 * time.Hour)

    // Create the JWT claims
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["user_id"] = username
    claims["exp"] = expirationTime

    // Generate the JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtSecret := []byte(configs.Cfg.JwtSecret)
    signedToken, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    
    return signedToken, nil
}

func GenerateRefreshToken(userID string) (string, error) {
    // Set the expiration time for the refresh token
    expirationTime := time.Now().Add(24 * time.Hour)
    
    // Create the JWT claims
    claims := jwt.MapClaims{}
    claims["authorized"] = true
    claims["user_id"] = userID
    claims["exp"] = expirationTime


    // Generate the JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    jwtSecret := []byte(configs.Cfg.JwtSecret)
    signedToken, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    
    return signedToken, nil
}