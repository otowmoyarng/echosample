package project

import (
	"net/http"

	"github.com/labstack/echo"
)

type Project struct {
	Team   string `json:"team"`
	Menber string `json:"menber"`
}

func Pathget(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func Queryget(c echo.Context) error {
	return c.JSON(http.StatusOK, Project{Team: c.QueryParam("team"), Menber: c.QueryParam("menber")})
}

func Bodyget(c echo.Context) error {

	result := new(Project)
	if err := c.Bind(result); err != nil {
		return err
	}

	result.Team += "hoge"
	result.Menber += "yamada"
	return c.JSON(http.StatusOK, result)
}
