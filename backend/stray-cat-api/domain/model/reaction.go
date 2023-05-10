package model

import "time"

type Reaction struct {
	ReactionID string
	UserID     string
	CatID      string
	DateTime   time.Time
}
