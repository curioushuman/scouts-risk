package model

import "time"

type Event struct {
	Label       string
	Id          string
	Description string
	Date 		    time.Time
	Location    Location
	Activities  []Activity
}
