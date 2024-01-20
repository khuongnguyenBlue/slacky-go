package services

import (
	"github.com/khuongnguyenBlue/slacky/models"
	"github.com/khuongnguyenBlue/slacky/transport"
)

type IAuthService interface {
	Register(registrationRequest transport.RegistrationRequest) (*models.User, error)
	Login(loginRequest transport.LoginRequest) (token string, err error)
}