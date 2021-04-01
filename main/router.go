package main

import (
	"echosample/api"
	"echosample/view"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const API_VERSION string = "1.1"

func NewRouter() *echo.Echo {

	entrypoint := fmt.Sprintf("/v%s", API_VERSION)

	e := echo.New()
	templaterender := &TemplateRenderer{
		templetes: template.Must(template.ParseGlob("../templete/*.html")),
	}
	e.Renderer = templaterender

	// ロギング
	log, err := os.OpenFile("../log/echosample.log", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		return e
	}
	// e.Use(middleware.Logger())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: log,
	}))

	// ルート
	e.GET(entrypoint, func(c echo.Context) error {
		data := fmt.Sprintf("echosample Ver.%s", API_VERSION)
		return c.String(http.StatusOK, data)
	})

	// API群
	api_users := e.Group(entrypoint + "/user")
	api_users.GET("/:name", api.GetUser)
	api_users.GET("", api.GetUsers)
	api_users.POST("", api.CreateUsers)
	api_users.PUT("/:name", api.UpdateUser)
	api_users.DELETE("/:name", api.DeleteUser)

	// Controller群
	userview := e.Group("/userview")
	userview.GET("", view.GetUsers)

	return e
}
