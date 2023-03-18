package main

import (
	"fmt"
	// "net/http"

	"golang-auth/configs"
	"golang-auth/db"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var e = echo.New()

func init(){
	err := cleanenv.ReadEnv(&configs.Cfg)
	fmt.Printf("%v", configs.Cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configuration")
	}
}

func main() {
	
	db.ConnectDB()
	
	// err := cleanenv.ReadEnv(&configs.Cfg)
	// if err != nil {
	// 	e.Logger.Fatal("Unable to load configuration")
	// }

	
	// e.POST("/save", func(c echo.Context) error {
	// 	name := c.FormValue("name")
	// 	email := c.FormValue("email")
	// 	return c.String(http.StatusOK, "name:" + name + ", email:" + email)
	// })

	// e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", configs.Cfg.Port)))
}
