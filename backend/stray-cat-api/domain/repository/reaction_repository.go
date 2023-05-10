// domain/repository/reaction_repository.go
package repository

import "stray-cat-api/domain/model"

type ReactionRepository interface {
	FindByID(reactionID int) (*model.Reaction, error)
	Store(reaction *model.Reaction) error
	Update(reaction *model.Reaction) error
	Delete(reactionID int) error
}
