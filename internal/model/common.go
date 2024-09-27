package model

type Tag string

// ! NOTE: Currently not used
// Still working out embeddings
type Base struct {
	Label string
	Id    string
	// Tags  []Tag
}

// ! Attempts at a union-ish tyoe
// type WithHazards struct {
// 	Base
// 	Hazards []Hazard
// }

// type WithHazards struct {
// 	Activity
// 	Location
// }

// type WithHazards interface {
// 	GetHazards() []Hazard
// }

// func (loc Location) GetHazards() {}
// func (act Activity) GetHazards() {}