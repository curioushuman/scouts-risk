package model

// Poor mans string enum
// ! NOTE: it does not prevent c.When = ActivityType("garbage")
type ActivityType string
const (
	ActivityTypeFormal   ActivityType = "formal"
	ActivityTypeInformal ActivityType = "informal"
)

type Activity struct {
	Label   string
	Id      string
	Hazards []Hazard
	ActType  ActivityType
}
