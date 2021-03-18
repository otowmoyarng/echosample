package main

import (
	"net/http"

	"echosample/api"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo")
	})

	api_users := e.Group("/user")
	api_users.GET("/:name", api.GetUser)
	api_users.GET("", api.GetUsers)
	api_users.POST("", api.CreateUsers)
	api_users.PUT("/:name", api.UpdateUser)
	api_users.DELETE("/:name", api.DeleteUser)
	return e
}
