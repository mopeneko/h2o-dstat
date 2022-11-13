package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mopeneko/h2o-dstat/backend/controllers"
)

func main() {
	e := echo.New()

	controllers.InitRouter(e)

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Logger.Fatal(e.Start(":3000"))
}
