package main

import (
	"go-rest-practice/db"
	"go-rest-practice/handler"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.InitDB()
	defer db.CloseDB()

	e := echo.New()

	u := e.Group("/users")
	u.Use(echojwt.JWT([]byte("secret")))
	u.GET("/:id", handler.GetUserDetail)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)

	e.Logger.Fatal(e.Start(":8080"))
}
