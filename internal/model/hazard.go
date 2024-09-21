package model

// Poor mans string enum
// ! NOTE: it does not prevent c.When = HazardWhen("garbage")
type HazardWhen string
const (
	HazardWhenBefore   HazardWhen = "before"
	HazardWhenDuring HazardWhen = "during"
	HazardWhenResponse  HazardWhen = "response"
)

type Hazard struct {
	Label       string
	Id		      string
	Controls []Control
	Hazards []Hazard
}
