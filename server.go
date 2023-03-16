package main

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.POST("/save", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, "name:" + name + ", email:" + email)
	})

e.Logger.Fatal(e.Start(":1323"))
}


