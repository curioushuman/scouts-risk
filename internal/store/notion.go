package store

import (
	"os"

	"github.com/jomei/notionapi"
)

// ! Eventually we don't want this to be exportable
// We should wrap this in a port interface
// and the Notion implementation should be in a different package
// in the form of an adapter
func NewNotionClient() *notionapi.Client {
	var notion_token = os.Getenv("STORE_NOTION_TOKEN")
	client := notionapi.NewClient(notionapi.Token(notion_token))
	return client
}