package model

import "time"

type Reaction struct {
	ReactionID int
	UserID     string
	CatID      string
	DateTime   time.Time
}
