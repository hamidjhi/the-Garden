package main

import (
	"chemex/config"
	"chemex/controller"
	"chemex/db"
	"github.com/labstack/echo"
)

func main(){
	db.AutoMigrate()
	defer db.MySQL.Close()
	e := echo.New()
	config.RegisterMiddlewares(e)
	controller.RegisterRoutes(e)
	e.Logger.Fatal(e.Start(":" + config.Port))
}
