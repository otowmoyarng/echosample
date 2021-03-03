package main

// type Project struct {
// 	Team   string `json:"team"`
// 	Menber string `json:"menber"`
// }

func main() {
	// e := echo.New()
	// e.GET("/index", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello echo")
	// })
	// e.GET("/index/:id", pathget)
	// e.GET("/project", queryget)
	// e.POST("/project", bodyget)
	// e.Logger.Fatal(e.Start(":1234"))
	router := NewRouter()
	router.Logger.Fatal(router.Start(":1234"))
}

// func pathget(c echo.Context) error {
// 	return c.String(http.StatusOK, c.Param("id"))
// }

// func queryget(c echo.Context) error {
// 	return c.JSON(http.StatusOK, Project{Team: c.QueryParam("team"), Menber: c.QueryParam("menber")})
// }

// func bodyget(c echo.Context) error {

// 	result := new(Project)
// 	if err := c.Bind(result); err != nil {
// 		return err
// 	}

// 	result.Team += "hoge"
// 	result.Menber += "yamada"
// 	return c.JSON(http.StatusOK, result)
// }
