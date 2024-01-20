//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/khuongnguyenBlue/slacky/controllers"
	"github.com/khuongnguyenBlue/slacky/infrastructures"
	"github.com/khuongnguyenBlue/slacky/interfaces"
	repoInterface "github.com/khuongnguyenBlue/slacky/interfaces/repositories"
	serviceInterface "github.com/khuongnguyenBlue/slacky/interfaces/services"
	"github.com/khuongnguyenBlue/slacky/repositories"
	"github.com/khuongnguyenBlue/slacky/services"
)

var controllersSet = wire.NewSet(
	controllers.NewAuthController,
)

var sericesSet = wire.NewSet(
	services.NewAuthService,
	wire.Bind(new(serviceInterface.IAuthService), new(*services.AuthService)),
)

var repositoriesSet = wire.NewSet(
	repositories.NewUserRepository,
	wire.Bind(new(repoInterface.IUserRepository), new(*repositories.UserRepository)),
	infrastructures.NewPgHandler,
	wire.Bind(new(interfaces.IDbHandler), new(*infrastructures.PgHandler)),
)

func newValidator() *validator.Validate {
	return validator.New(validator.WithRequiredStructEnabled())
}

func initializeApp() *app {
	wire.Build(controllersSet, sericesSet, repositoriesSet, newValidator, wire.Struct(new(app), "*"))
	return &app{}
}
