// usecase/interactor/stray_cat_interactor.go
package interactor

import (
	"stray-cat-api/domain/model"
	"stray-cat-api/domain/repository"
)

type StrayCatInteractor struct {
	StrayCatRepository repository.StrayCatRepository
}

func (i *StrayCatInteractor) FindAll() ([]*model.StrayCat, error) {
	return i.StrayCatRepository.FindAll()
}

func (i *StrayCatInteractor) FindByID(catID int) (*model.StrayCat, error) {
	return i.StrayCatRepository.FindByID(catID)
}

func (i *StrayCatInteractor) Store(strayCat *model.StrayCat) error {
	return i.StrayCatRepository.Store(strayCat)
}

func (i *StrayCatInteractor) Update(strayCat *model.StrayCat) error {
	return i.StrayCatRepository.Update(strayCat)
}

func (i *StrayCatInteractor) Delete(catID int) error {
	return i.StrayCatRepository.Delete(catID)
}
