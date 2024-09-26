package fixtures

import (
	"time"

	"scouts-risk/internal/model"
)

var Events = map[string]model.Event{
	"campAtMacs": {
		Label: "Campout at Mac's Place",
  	Id: "campAtMacs",
		Description: "A weekend campout at Mac's Place.",
		Date: time.Date(2024, time.September, 7, 0, 0, 0, 0, time.UTC),
		Location: Locations["macsPlace"],
	},
	"anotherCampAtMacs": {
		Label: "Another Campout at Mac's Place",
  	Id: "anotherCampAtMacs",
		Description: "ANOTHER weekend campout at Mac's Place.",
		Date: time.Date(2024, time.November, 7, 0, 0, 0, 0, time.UTC),
		Location: Locations["macsPlace"],
	},
	"campAtUlladulla": {
		Label: "Campout at Ulladulla",
  	Id: "campAtUlladulla",
		Description: "Bike camp round 2",
		Date: time.Date(2024, time.December, 7, 0, 0, 0, 0, time.UTC),
		Location: LocationLocationEmpty,
	},
}