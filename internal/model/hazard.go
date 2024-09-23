package model

type Hazard struct {
	Label    string
	Id		   string
	Controls []Control
	// Tags  []Tag
}

type Consequence struct {
	Label    string
	Id		   string
	// Tags  []Tag
}