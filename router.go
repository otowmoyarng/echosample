package main

import (
	"echosample/project"
	"net/http"

	"github.com/labstack/echo"
)

func NewRouter() *echo.Echo {
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo")
	})

	apipro := e.Group("/project")
	apipro.GET("/:id", project.Pathget)
	apipro.GET("", project.Queryget)
	apipro.POST("", project.Bodyget)
	return e
}
