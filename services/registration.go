package services

import (
	"github.com/khuongnguyenBlue/slacky/interfaces/repositories"
	"github.com/khuongnguyenBlue/slacky/models"
	"github.com/khuongnguyenBlue/slacky/transport"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	repositories.IUserRepository
}

func NewRegistrationService(userRepo repositories.IUserRepository) *RegistrationService {
	return &RegistrationService{userRepo}
}

func (service *RegistrationService) Register(request transport.RegistrationRequest) (*models.User, error) {
	hash, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user, err := service.CreateUser(request.Email, hash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
