// domain/repository/user_info_repository.go
package repository

import "stray-cat-api/domain/model"

type UserInfoRepository interface {
	FindAll() ([]*model.UserInfo, error)
	FindByID(userID string) (*model.UserInfo, error)
	Store(user *model.UserInfo) error
	Update(user *model.UserInfo) error
	Delete(userID string) error
}
