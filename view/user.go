package view

import (
	"echosample/model"
	"net/http"

	"github.com/labstack/echo"
)

func GetUsers(c echo.Context) error {
	userlist := model.GetUsers()
	return c.Render(http.StatusOK, "index.html", userlist)
}
