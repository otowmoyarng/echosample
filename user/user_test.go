package user

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testvalue string = "hoge"
	c.SetParamNames("id")
	c.SetParamValues(testvalue)

	GetUser(c)
	fmt.Printf("HTTPStatus:%d\n", rec.Code)
	fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, testvalue, rec.Body.String())
}

func TestGetUsers(t *testing.T) {

	queryparams := make(url.Values)
	queryparams.Set("team", "hoge")
	queryparams.Set("menber", "fuga")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/user/?"+queryparams.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testJSON string = `{"team":"hoge","menber":"fuga"}`
	GetUsers(c)
	fmt.Printf("HTTPStatus:%d\n", rec.Code)
	fmt.Printf("json=%s\n", testJSON)
	fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, testJSON+"\n", rec.Body.String())
}

func TestCreateUsers(t *testing.T) {

	var testParamJSON string = `{"team":"hoge","menber":"fuga"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/user/", strings.NewReader(testParamJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	CreateUsers(c)

	var testExpectJSON string = `{"team":"hogehoge","menber":"fugafuga"}`
	fmt.Printf("HTTPStatus:%d\n", rec.Code)
	fmt.Printf("json=%s\n", testParamJSON)
	fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, testExpectJSON+"\n", rec.Body.String())
}
