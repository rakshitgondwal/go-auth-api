package main

import (
	"fmt"
	"net/http"

	"golang-auth/configs"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	
	err := cleanenv.ReadEnv(&configs.Cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}

	
	e.POST("/save", func(c echo.Context) error {
		name := c.FormValue("name")
		email := c.FormValue("email")
		return c.String(http.StatusOK, "name:" + name + ", email:" + email)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", configs.Cfg.Port)))
}
