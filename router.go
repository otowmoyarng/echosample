package main

import (
	"echosample/user"
	"net/http"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo")
	})

	apipro := e.Group("/user")
	apipro.GET("/:id", user.GetUser)
	apipro.GET("", user.GetUsers)
	apipro.POST("", user.CreateUsers)
	return e
}
