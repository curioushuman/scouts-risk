package api_location

import (
	"reflect"

	"github.com/labstack/echo/v4"

	"scouts-risk/internal/fixtures"
	"scouts-risk/internal/model"
	"scouts-risk/internal/util"
	web_location "scouts-risk/web/location"
)

func New(c echo.Context) error {
	label := c.FormValue("label")
	locType := c.FormValue("type")
	templateId := c.FormValue("template")

	newItem := model.Location{}
	if (templateId != "") {
		existingItem := util.FindByStringId(templateId, fixtures.LocTypeLocations(model.LocationType(locType)))
		if (!reflect.ValueOf(existingItem).IsZero()) {
			newItem = existingItem
		}
  }
 	newItem.Label = label
	newItem.LocType = model.LocationType(locType)

	html := web_location.ReviewLocation(newItem)
	return html.Render(c.Request().Context(), c.Response().Writer)
}
