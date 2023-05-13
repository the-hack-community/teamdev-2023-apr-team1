package model

import "time"

type Reaction struct {
	ReactionID int       `json:"reactionId"`
	UserID     string    `json:"userId"`
	CatID      string    `json:"catId"`
	DateTime   time.Time `json:"dateTime"`
}
