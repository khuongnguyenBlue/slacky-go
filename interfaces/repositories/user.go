package repositories

import "github.com/khuongnguyenBlue/slacky/models"

type IUserRepository interface {
	CreateUser(email string, hash string) (*models.User, error)
}