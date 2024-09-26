package api_location

import (
	"github.com/labstack/echo/v4"

	"scouts-risk/internal/fixtures"
	"scouts-risk/internal/model"
	web_location "scouts-risk/web/location"
)

func New(c echo.Context) error {
	newLabel := c.FormValue("label")
	newType := c.FormValue("type")
	newTemplate := c.FormValue("template")

	newArea := fixtures.LocationArea
	newArea.Label = newLabel
	newArea.LocType = model.LocationType(newType)
	if (newTemplate == "") {
		newArea.Hazards = []model.Hazard{}
	}

	html := web_location.ReviewLocation(newArea)
	return html.Render(c.Request().Context(), c.Response().Writer)
}