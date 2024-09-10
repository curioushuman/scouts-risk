package api

import (
	"github.com/labstack/echo/v4"
)

func PostTst(c echo.Context) error {
	q := c.FormValue("question")
	html := TstHtmx(q)
	return html.Render(c.Request().Context(), c.Response().Writer)
}