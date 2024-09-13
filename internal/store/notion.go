package store

import (
	"log"
	"os"

	"github.com/jomei/notionapi"
)

// Local implementation of the Store interface
type service struct {
	store *notionapi.Client
}

var (
	token = os.Getenv("STORE_NOTION_TOKEN")
	storeService *service
)

func New() Service {
	// Reuse service if exists
	if storeService != nil {
		return storeService
	}

	// TODO I believe this is going to `panic`
	// we'll need to handle this error
  	store := notionapi.NewClient(notionapi.Token(token))

	storeService = &service{
		store: store,
	}
	return storeService
}

// Close closes the connection.
func (s *service) Close() error {
	log.Printf("There is no need to disconnect from Notion")
	return nil
}