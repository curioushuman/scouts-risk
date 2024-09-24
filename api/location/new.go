package api_location

import (
	"github.com/labstack/echo/v4"

	"scouts-risk/internal/fixtures"
	web_location "scouts-risk/web/location"
)

func New(c echo.Context) error {
	newLabel := c.FormValue("label")
	newArea := fixtures.LocationArea
	newArea.Label = newLabel

	html := web_location.Review(newArea)
	return html.Render(c.Request().Context(), c.Response().Writer)
}