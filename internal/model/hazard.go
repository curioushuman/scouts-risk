package model

type Hazard struct {
	Label 	 string
	Id    	 string
	Controls []Control
}

type Consequence struct {
	Label string
	Id    string
}
