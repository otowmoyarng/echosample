package api

import (
	repo "echosample/repository"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUser_found(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:name", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testvalue string = "taro"
	c.SetParamNames("name")
	c.SetParamValues(testvalue)

	GetUser(c)

	expect_json, _ := json.Marshal(repo.User{Name: testvalue, Birthday: "2010/01/01"})
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(expect_json)+"\n", rec.Body.String())
}

func TestGetUser_notfound(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:name", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testvalue string = "maki"
	c.SetParamNames("name")
	c.SetParamValues(testvalue)

	GetUser(c)

	assert.Equal(t, http.StatusNotFound, rec.Code)
}
func TestGetUsers(t *testing.T) {

	// queryparams := make(url.Values)
	// queryparams.Set("team", "hoge")
	// queryparams.Set("menber", "fuga")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/", nil)
	// req := httptest.NewRequest(http.MethodGet, "/user/?"+queryparams.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// var testJSON string = `{"team":"hoge","menber":"fuga"}`
	GetUsers(c)
	// fmt.Printf("HTTPStatus:%d\n", rec.Code)
	// fmt.Printf("json=%s\n", testJSON)
	// fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	// assert.Equal(t, testJSON+"\n", rec.Body.String())
}

func TestCreateUsers(t *testing.T) {

	paramuser_json, _ := json.Marshal(repo.User{Name: "koji", Birthday: "1990/12/31"})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/", strings.NewReader(string(paramuser_json)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUsers(c)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "create success!", rec.Body.String())
}

func TestUpdateUser(t *testing.T) {

	paramuser_json, _ := json.Marshal(repo.User{Name: "koji", Birthday: "1995/10/13"})

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/user/:name", strings.NewReader(string(paramuser_json)))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testvalue string = "koji"
	c.SetParamNames("name")
	c.SetParamValues(testvalue)

	UpdateUser(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "update success!", rec.Body.String())
}

func TestDeleteUser(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/user/:name", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testvalue string = "koji"
	c.SetParamNames("name")
	c.SetParamValues(testvalue)

	DeleteUser(c)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, "delete success!", rec.Body.String())
}
