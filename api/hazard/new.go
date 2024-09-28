package api_hazard

import (
	"reflect"

	"github.com/labstack/echo/v4"

	"scouts-risk/internal/fixtures"
	"scouts-risk/internal/model"
	"scouts-risk/internal/util"
	web_hazard "scouts-risk/web/hazard"
)

func New(c echo.Context) error {
	label := c.FormValue("label")
	existingStringId := c.FormValue("existing")
	parentId := c.FormValue("parentId")

	newItem := model.Hazard{}
	newItem.Label = label
	if (existingStringId != "") {
		 existingHazard := util.FindByStringId(existingStringId, fixtures.HazardsFirstAid)
		 if (!reflect.ValueOf(existingHazard).IsZero()) {
			 newItem = existingHazard
		 }
	}

	html := web_hazard.ReviewHazard(newItem, parentId)
	return html.Render(c.Request().Context(), c.Response().Writer)
}

func NewControl(c echo.Context) error {
	label := c.FormValue("label")
	existingStringId := c.FormValue("existing")
	// parentId := c.FormValue("parentId")

	newItem := model.Control{}
	newItem.Label = label
	if (existingStringId != "") {
		 existingHazard := util.FindByStringId(existingStringId, fixtures.ControlsExisting)
		 if (!reflect.ValueOf(existingHazard).IsZero()) {
			 newItem = existingHazard
		 }
	}

	html := web_hazard.ReviewControl(newItem)
	return html.Render(c.Request().Context(), c.Response().Writer)
}

func NewConsequence(c echo.Context) error {
	label := c.FormValue("label")
	existingStringId := c.FormValue("existing")
	// parentId := c.FormValue("parentId")

	newItem := model.Consequence{}
	newItem.Label = label
	if (existingStringId != "") {
		 existingHazard := util.FindByStringId(existingStringId, fixtures.ConsequencesExisting)
		 if (!reflect.ValueOf(existingHazard).IsZero()) {
			 newItem = existingHazard
		 }
	}

	html := web_hazard.ReviewConsequence(newItem)
	return html.Render(c.Request().Context(), c.Response().Writer)
}
