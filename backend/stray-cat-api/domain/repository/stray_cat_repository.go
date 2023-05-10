// domain/repository/stray_cat_repository.go
package repository

import "stray-cat-api/domain/model"

type StrayCatRepository interface {
	FindAll() ([]*model.StrayCat, error)
	FindByID(catID string) (*model.StrayCat, error)
	Store(strayCat *model.StrayCat) error
	Update(strayCat *model.StrayCat) error
	Delete(catID string) error
}
