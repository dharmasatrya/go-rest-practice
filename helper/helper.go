package helper

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GetClaimsFromToken(c echo.Context) (jwt.MapClaims, error) {
	// Extract the token from the context
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}

	// Extract claims from the token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}

	// Return the claims
	return claims, nil
}
