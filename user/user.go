package user

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Team   string `json:"team"`
	Menber string `json:"menber"`
}

func GetUser(c echo.Context) error {
	return c.String(http.StatusOK, c.Param("id"))
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, User{Team: c.QueryParam("team"), Menber: c.QueryParam("menber")})
}

func CreateUsers(c echo.Context) error {

	result := new(User)
	if err := c.Bind(result); err != nil {
		return err
	}

	result.Team += "hoge"
	result.Menber += "fuga"
	return c.JSON(http.StatusOK, result)
}
