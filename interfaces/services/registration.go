package services

import (
	"github.com/khuongnguyenBlue/slacky/models"
	"github.com/khuongnguyenBlue/slacky/transport"
)

type IRegistrationService interface {
	Register(registrationRequest transport.RegistrationRequest) (*models.User, error)
}