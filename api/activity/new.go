package api_activity

import (
	"reflect"

	"github.com/labstack/echo/v4"

	"scouts-risk/internal/fixtures"
	"scouts-risk/internal/model"
	"scouts-risk/internal/util"
	web_activity "scouts-risk/web/activity"
)

func New(c echo.Context) error {
	label := c.FormValue("label")
	actType := c.FormValue("type")
	duplicateId := c.FormValue("duplicate")

	newItem := model.Activity{}
	if (duplicateId != "") {
		existingItem := util.FindByStringId(duplicateId, fixtures.ActTypeActivities(model.ActivityType(actType)))
		if (!reflect.ValueOf(existingItem).IsZero()) {
			newItem = existingItem
		}
  }
 	newItem.Label = label
	newItem.ActType = model.ActivityType(actType)

	html := web_activity.ReviewActivity(newItem)
	return html.Render(c.Request().Context(), c.Response().Writer)
}
