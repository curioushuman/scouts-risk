package api

import (
	"scouts-risk/internal/store"

	"github.com/labstack/echo/v4"
)

func PostTst(c echo.Context) error {
	// q := c.FormValue("question")
	// https://www.notion.so/curioushuman/Campout-at-Mac-s-c77982b36df84bc3a735dace06b66d71?pvs=4
	// c77982b36df84bc3a735dace06b66d71
	client := store.NewNotionClient()
	page, err := client.Page.Get(c.Request().Context(), "c77982b36df84bc3a735dace06b66d71")
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	html := TstHtmx(page.Properties["Name"].GetID())
	return html.Render(c.Request().Context(), c.Response().Writer)
}