// usecase/interactor/reaction_interactor.go
package interactor

import (
	"stray-cat-api/domain/model"
	"stray-cat-api/domain/repository"
)

type ReactionInteractor struct {
	ReactionRepository repository.ReactionRepository
}

func (i *ReactionInteractor) FindByID(reactionID string) (*model.Reaction, error) {
	return i.ReactionRepository.FindByID(reactionID)
}

func (i *ReactionInteractor) Store(reaction *model.Reaction) error {
	return i.ReactionRepository.Store(reaction)
}

func (i *ReactionInteractor) Update(reaction *model.Reaction) error {
	return i.ReactionRepository.Update(reaction)
}

func (i *ReactionInteractor) Delete(reactionID string) error {
	return i.ReactionRepository.Delete(reactionID)
}
