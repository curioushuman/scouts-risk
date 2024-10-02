package api_feedback

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"scouts-risk/internal/ui"

	"github.com/jomei/notionapi"
	"github.com/labstack/echo/v4"
)

func Basic(c echo.Context) error {
	// save feedback
	// client := store.New()
	// ! Skipping store for now, we'll cross that bridge later
	token := os.Getenv("STORE_NOTION_TOKEN")
	client := notionapi.NewClient(notionapi.Token(token))
	_, err := client.Page.Create(c.Request().Context(), prepareNewPage(c))
	if err != nil {
		log.Printf("Error creating feedback: %v", err)
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	// return success as toast
	html := ui.Toast(map[string]string{
		"success": "Your feedback has been received, thank you!",
		})
	return html.Render(c.Request().Context(), c.Response().Writer)
}

func prepareNewPage(c echo.Context) *notionapi.PageCreateRequest {
	// database Id will be the same for all feedback
	// https://www.notion.so/curioushuman/11258db6467380f3ba40c0ac855f5aad?v=e2f12bd325454c90ac37d43f848924ea&pvs=4
	req := &notionapi.PageCreateRequest{
		Parent: notionapi.Parent{
			Type:       notionapi.ParentTypeDatabaseID,
			DatabaseID: "11258db6467380f3ba40c0ac855f5aad",
		},
		Properties: notionapi.Properties{
			"Name": notionapi.TitleProperty{
				Title: []notionapi.RichText{
					{Text: &notionapi.Text{Content: time.Now().Format("2006-01-02 15:04:05")}},
				},
			},
			"Identifier": textProperty("uniqueId", c),
			"Specific": textProperty("specific", c),
			"Feedback": textProperty("feedback", c),
			"OS": textProperty("os", c),
			"Browser": textProperty("browser", c),
			"Use": numberProperty("use", c),
			"Recommend": numberProperty("recommend", c),
		},
	}
	return req
}

func textProperty(key string, c echo.Context) (p notionapi.RichTextProperty) {
	log.Printf("formValue(%v): %v", key, c.FormValue(key))
	p = notionapi.RichTextProperty{
		Type: notionapi.PropertyTypeRichText,
		RichText: []notionapi.RichText{
			{
				Type: notionapi.ObjectTypeText,
				Text: &notionapi.Text{Content: c.FormValue(key)},
			},
		},
	}
	return
}

func numberProperty(key string, c echo.Context) (p notionapi.NumberProperty) {
	fl, _ := strconv.ParseFloat(strings.TrimSpace(c.FormValue(key)), 64)

	p = notionapi.NumberProperty{
		Type: notionapi.PropertyTypeNumber,
		Number: fl,
	}
	return
}