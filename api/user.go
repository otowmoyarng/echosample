package api

import (
	"echosample/model"
	repo "echosample/repository"
	"net/http"

	"github.com/labstack/echo"
)

func getCondition(c echo.Context) repo.User {
	condition := new(repo.User)
	condition.Name = c.Param("name")
	return *condition
}

func GetUser(c echo.Context) error {
	condition := getCondition(c)
	result, isok := model.GetUser(condition)
	if isok != nil {
		return echo.NewHTTPError(http.StatusNotFound, isok.Error())
	}
	return c.JSON(http.StatusOK, result)
}

func GetUsers(c echo.Context) error {
	userlist := model.GetUsers()
	return c.JSON(http.StatusOK, userlist)
}

func CreateUsers(c echo.Context) error {

	insertuser := new(repo.User)
	if err := c.Bind(insertuser); err != nil {
		return err
	}
	model.InsertUser(*insertuser)
	return c.String(http.StatusOK, "create success!")
}

func UpdateUser(c echo.Context) error {

	updatejson := new(repo.User)
	if err := c.Bind(updatejson); err != nil {
		return err
	}
	updateuser := getCondition(c)
	updateuser.Birthday = updatejson.Birthday
	model.UpdateUser(updateuser)
	return c.String(http.StatusOK, "update success!")
}

func DeleteUser(c echo.Context) error {

	deleteuser := getCondition(c)
	model.DeleteUser(deleteuser)
	return c.String(http.StatusOK, "delete success!")
}
