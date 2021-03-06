package main

import (
	"GO-REGRIS-LOGIN/config"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func IsLoggedIn() echo.MiddlewareFunc {
	cfg, _ := config.NewConfig(".env")
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(cfg.JWTConfig.SecretKey),
	})
}
