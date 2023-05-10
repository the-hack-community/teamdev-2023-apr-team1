// domain/repository/reaction_repository.go
package repository

import "stray-cat-api/domain/model"

type ReactionRepository interface {
	FindByID(reactionID string) (*model.Reaction, error)
	Store(reaction *model.Reaction) error
	Update(reaction *model.Reaction) error
	Delete(reactionID string) error
}
