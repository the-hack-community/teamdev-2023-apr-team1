package model

import "time"

type StrayCat struct {
	CatID           int        `json:"catId"`
	UserID          string     `json:"userId"`
	PhotoData       string     `json:"photoData"`
	CaptureDateTime time.Time  `json:"captureDateTime"`
	Location        Location   `json:"location"`
	Name            string     `json:"name"`
	Features        string     `json:"features"`
	Condition       string     `json:"condition"`
	Reactions       []Reaction `json:"reactions"`
}
