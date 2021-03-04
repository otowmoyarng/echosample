package project

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

func TestPathget(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/project/:id", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testvalue string = "hoge"
	c.SetParamNames("id")
	c.SetParamValues(testvalue)

	Pathget(c)
	fmt.Printf("HTTPStatus:%d\n", rec.Code)
	fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, testvalue, rec.Body.String())
}

func TestQueryget(t *testing.T) {

	queryparams := make(url.Values)
	queryparams.Set("team", "hoge")
	queryparams.Set("menber", "fuga")

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/project/?"+queryparams.Encode(), nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	var testJSON string = `{"team":"hoge","menber":"fuga"}`
	Queryget(c)
	fmt.Printf("HTTPStatus:%d\n", rec.Code)
	fmt.Printf("json=%s\n", testJSON)
	fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, testJSON+"\n", rec.Body.String())
}

func TestBodyget(t *testing.T) {

	var testParamJSON string = `{"team":"hoge","menber":"fuga"}`

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/project/", strings.NewReader(testParamJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	Bodyget(c)

	var testExpectJSON string = `{"team":"hogehoge","menber":"fugafuga"}`
	fmt.Printf("HTTPStatus:%d\n", rec.Code)
	fmt.Printf("json=%s\n", testParamJSON)
	fmt.Printf("result=%s\n", rec.Body.String())
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, testExpectJSON+"\n", rec.Body.String())
}
