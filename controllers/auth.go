package controllers

import (
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	interfaces "github.com/khuongnguyenBlue/slacky/interfaces/services"
	"github.com/khuongnguyenBlue/slacky/transport"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	registrationService interfaces.IRegistrationService
	validator		   *validator.Validate
}

func NewAuthController(registrationService interfaces.IRegistrationService, validator *validator.Validate) *AuthController {
	return &AuthController{registrationService, validator}
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

	user, err := controller.registrationService.Register(*registrationRequest)
	if err != nil {
		if pgError, ok := err.(*pgconn.PgError); ok {
			if pgError.Code == "23505" {
				return echo.NewHTTPError(400, "email already exists")
			}
		}

		return err
	}

	return c.JSON(200, transport.SuccessResponse(transport.ToRegistrationData(*user)))
}