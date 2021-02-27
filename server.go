package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type Project struct {
	Team   string `json:"team"`
	Menber string `json:"menber"`
}

func main() {
	e := echo.New()
	e.GET("/index", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello echo")
	})
	e.GET("/index/:id", pathget)
	e.GET("/project", queryget)
	e.POST("/project", bodyget)
	e.Logger.Fatal(e.Start(":1234"))
}

func pathget(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func queryget(c echo.Context) error {

	result := new(Project)
	result.Team = c.QueryParam("team")
	result.Menber = c.QueryParam("menber")
	return c.JSON(http.StatusOK, result)
}

func bodyget(c echo.Context) error {

	result := new(Project)
	if err := c.Bind(result); err != nil {
		return err
	}

	result.Team += "hoge"
	result.Menber += "yamada"
	return c.JSON(http.StatusOK, result)
}
