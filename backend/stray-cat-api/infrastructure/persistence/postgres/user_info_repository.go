package postgres

import (
	"database/sql"
	"stray-cat-api/domain/model"
	"stray-cat-api/domain/repository"
)

type UserInfoRepository struct {
	DB *sql.DB
}

func NewUserInfoRepository(db *sql.DB) repository.UserInfoRepository {
	return &UserInfoRepository{DB: db}
}

func (r *UserInfoRepository) FindAll() ([]*model.UserInfo, error) {
	rows, err := r.DB.Query("SELECT user_id, name, email FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]*model.UserInfo, 0)
	for rows.Next() {
		user := new(model.UserInfo)
		err := rows.Scan(&user.UserID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserInfoRepository) FindByID(userId string) (*model.UserInfo, error) {
	user := new(model.UserInfo)
	err := r.DB.QueryRow("SELECT id, name, email FROM users WHERE user_id = $1", userId).Scan(&user.UserID, &user.Name, &user.Email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserInfoRepository) Store(user *model.UserInfo) error {
	err := r.DB.QueryRow("INSERT INTO users(name, email) VALUES($1, $2) RETURNING user_id", user.Name, user.Email).Scan(&user.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserInfoRepository) Update(user *model.UserInfo) error {
	_, err := r.DB.Exec("UPDATE users SET name = $1, email = $2 WHERE user_id = $3", user.Name, user.Email, user.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserInfoRepository) Delete(userId string) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE user_id = $1", userId)
	if err != nil {
		return err
	}
	return nil
}
