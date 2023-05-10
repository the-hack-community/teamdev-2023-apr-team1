package interactor

import (
	"stray-cat-api/domain/model"
	"stray-cat-api/domain/repository"
)

type UserInfoInteractor struct {
	UserInfoRepository repository.UserInfoRepository
}

func NewUserInfoInteractor(userInfoRepository repository.UserInfoRepository) *UserInfoInteractor {
	return &UserInfoInteractor{UserInfoRepository: userInfoRepository}
}

func (i *UserInfoInteractor) FindAll() ([]*model.UserInfo, error) {
	return i.UserInfoRepository.FindAll()
}

func (i *UserInfoInteractor) FindByID(userId string) (*model.UserInfo, error) {
	return i.UserInfoRepository.FindByID(userId)
}

func (i *UserInfoInteractor) Store(user *model.UserInfo) error {
	return i.UserInfoRepository.Store(user)
}

func (i *UserInfoInteractor) Update(user *model.UserInfo) error {
	return i.UserInfoRepository.Update(user)
}

func (i *UserInfoInteractor) Delete(userId string) error {
	return i.UserInfoRepository.Delete(userId)
}
