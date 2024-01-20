package controllers

import (
	"github.com/go-playground/validator/v10"
	interfaces "github.com/khuongnguyenBlue/slacky/interfaces/services"
	"github.com/khuongnguyenBlue/slacky/transport"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	authService interfaces.IAuthService
	validator		   *validator.Validate
}

func NewAuthController(authService interfaces.IAuthService, validator *validator.Validate) *AuthController {
	return &AuthController{authService, validator}
}

func (controller *AuthController) Register(c echo.Context) (error) {
	registrationRequest := new(transport.RegistrationRequest)
	if err := c.Bind(registrationRequest); err != nil {
		return err
	}

	err := controller.validator.Struct(registrationRequest)
	if err != nil {
		return err
	}

	user, err := controller.authService.Register(*registrationRequest)
	if err != nil {
		return err
	}

	return c.JSON(200, transport.SuccessResponse(transport.ToRegistrationData(*user)))
}

func (controller *AuthController) Login(c echo.Context) (error) {
	loginRequest := new(transport.LoginRequest)
	if err := c.Bind(loginRequest); err != nil {
		return err
	}

	err := controller.validator.Struct(loginRequest)
	if err != nil {
		return err
	}

	token, err := controller.authService.Login(*loginRequest)
	if err != nil {
		return err
	}

	return c.JSON(200, transport.SuccessResponse(transport.LoginData{Token: token}))
}