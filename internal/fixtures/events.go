package fixtures

import (
	"time"

	"scouts-risk/internal/model"
)

var Events = map[string]model.Event{
	"campAtMacs": {
		Label: "Campout at Mac's Place",
  	Id: "campAtMacs",
		Date: time.Date(2024, time.September, 7, 0, 0, 0, 0, time.UTC),
		Locations: []model.Location{
			Locations["macsPlace"],
		},
	},
}