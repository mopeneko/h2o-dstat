package controllers

import (
	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/stats", GetStatus)
}
