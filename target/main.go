package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func postLoginHandler(c echo.Context) error {
	fmt.Println("login")
	user := c.Param("user")
	pass := c.FormValue("pass")
	fmt.Printf("update pass: (%s:%s)\n", user, pass)
	return c.String(http.StatusOK, "login")
}

func getLoginHandler(c echo.Context) error {
	user := c.QueryParam("user")
	pass := c.QueryParam("pass")

	if len(user) == 0 || len(pass) == 0 {
		return c.String(http.StatusNotFound, "login failed")
	}
	if user != "hoge" {
		return c.String(http.StatusNotFound, "user not found")
	}

	fmt.Printf("login: (%s:%s)\n", user, pass)
	fmt.Printf("login: (%s:%s)\n", user, pass)
	return c.String(http.StatusOK, "login")
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Target!")
	})
	e.GET("/login", getLoginHandler)
	e.POST("/login/:user", postLoginHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
