// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/khuongnguyenBlue/slacky/controllers"
	"github.com/khuongnguyenBlue/slacky/infrastructures"
	"github.com/khuongnguyenBlue/slacky/interfaces"
	repositories2 "github.com/khuongnguyenBlue/slacky/interfaces/repositories"
	services2 "github.com/khuongnguyenBlue/slacky/interfaces/services"
	"github.com/khuongnguyenBlue/slacky/repositories"
	"github.com/khuongnguyenBlue/slacky/services"
)

// Injectors from wire.go:

func initializeApp() *app {
	pgHandler := infrastructures.NewPgHandler()
	userRepository := repositories.NewUserRepository(pgHandler)
	registrationService := services.NewRegistrationService(userRepository)
	validate := newValidator()
	authController := controllers.NewAuthController(registrationService, validate)
	mainApp := &app{
		authController: authController,
	}
	return mainApp
}

// wire.go:

var controllersSet = wire.NewSet(controllers.NewAuthController)

var sericesSet = wire.NewSet(services.NewRegistrationService, wire.Bind(new(services2.IRegistrationService), new(*services.RegistrationService)))

var repositoriesSet = wire.NewSet(repositories.NewUserRepository, wire.Bind(new(repositories2.IUserRepository), new(*repositories.UserRepository)), infrastructures.NewPgHandler, wire.Bind(new(interfaces.IDbHandler), new(*infrastructures.PgHandler)))

func newValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}