package model

// Poor mans string enum
// ! NOTE: it does not prevent c.When = LocationType("garbage")
type LocationType string
const (
	LocationTypeLocation  LocationType = "location"
	LocationTypeArea     LocationType = "area"
	LocationTypeEquipment LocationType = "equipment"
)

type Location struct {
	Label    string
	Id		   string
	LocType  LocationType
	Hazards  []Hazard
	// Tags     []Tag
}
