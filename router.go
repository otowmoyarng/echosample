package main

import (
	"net/http"

	"echosample/api"

	"github.com/labstack/echo"
)

const api_version string = "v1.1"

func NewRouter() *echo.Echo {

	entrypoint := "/" + api_version
	e := echo.New()

	e.GET(entrypoint, func(c echo.Context) error {
		return c.String(http.StatusOK, "echosample Ver.1.1")
	})

	api_users := e.Group(entrypoint + "/user")
	api_users.GET("/:name", api.GetUser)
	api_users.GET("", api.GetUsers)
	api_users.POST("", api.CreateUsers)
	api_users.PUT("/:name", api.UpdateUser)
	api_users.DELETE("/:name", api.DeleteUser)
	return e
}
