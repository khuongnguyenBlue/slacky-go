package repositories

import (
	"github.com/khuongnguyenBlue/slacky/interfaces"
	"github.com/khuongnguyenBlue/slacky/models"
)

type UserRepository struct {
	interfaces.IDbHandler
}

func NewUserRepository(dbHandler interfaces.IDbHandler) *UserRepository {
	return &UserRepository{dbHandler}
}

func (repository *UserRepository) CreateUser(email string, hash string) (*models.User, error) {
	user := models.User{Email: email, Hash: hash}
	err := repository.Create(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
