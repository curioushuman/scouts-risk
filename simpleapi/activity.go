package simpleapi

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}
	return c.JSON(http.StatusOK, resp)
}
