package model

import "time"

type Event struct {
	Label      string
	Id         string
	Date 		   time.Time
	Locations  []Location
	Activities []Activity
}
