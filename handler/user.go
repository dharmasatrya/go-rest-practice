package handler

import (
	"context"
	"fmt"
	"go-rest-practice/db"
	"go-rest-practice/helper"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUser struct {
	Username string `json:"username" validate:"required, username"`
	Password string `json:"password" validate:"required"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required, username"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Register(c echo.Context) error {
	var req RegisterUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Request"})
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Invalid generate password"})
	}

	query := "INSERT INTO users(username, password) VALUES ($1, $2) RETURNING id, username"
	var userID int
	var userName string
	err = db.Pool.QueryRow(context.Background(), query, req.Username, string(hashPassword)).Scan(&userID, &userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Inserting user error"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "User Registered",
	})
}

var jwtSecret = []byte("secret")

func Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad request"})
	}

	var user User
	query := "SELECT id, username, password FROM users WHERE username = $1"
	err := db.Pool.QueryRow(context.Background(), query, req.Username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password1"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Email or Password 2"})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Error generating token"})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}

func GetUserDetail(c echo.Context) error {
	claims, err := helper.GetClaimsFromToken(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": err.Error()})
	}

	var user User
	query := "SELECT id, username FROM users WHERE id = $1"
	err = db.Pool.QueryRow(context.Background(), query, claims["user_id"]).Scan(&user.ID, &user.Username)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error fetching data"})
	}

	return c.JSON(http.StatusOK, user)
}
