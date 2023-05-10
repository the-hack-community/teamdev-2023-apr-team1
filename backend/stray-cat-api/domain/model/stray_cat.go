package model

import "time"

type StrayCat struct {
	CatID           int
	UserID          string
	PhotoData       string
	CaptureDateTime time.Time
	Location        Location
	Name            string
	Features        string
	Condition       string
	Reactions       []Reaction
}
