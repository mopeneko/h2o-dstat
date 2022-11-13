package echoutils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ReturnInternalServerError(c echo.Context) error {
	return c.String(http.StatusInternalServerError, "internal server error")
}
