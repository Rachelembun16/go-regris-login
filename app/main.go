package main

import (
	"GO-REGRIS-LOGIN/config"
	"GO-REGRIS-LOGIN/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	ech := echo.New()

	//Use route
	routes.InitRoute(ech)

	//Config
	cfg, _ := config.NewConfig(".env")

	ech.Logger.Fatal(ech.Start(":" + cfg.Port))
}
