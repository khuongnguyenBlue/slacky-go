package middlewares

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/khuongnguyenBlue/slacky/transport"
	"github.com/labstack/echo/v4"
)

// error handler middleware
type FieldValidationError struct {
	err validator.FieldError
}

func (fe FieldValidationError) Error() string {
	var sb strings.Builder

	if fe.err.ActualTag() == "required" {
		sb.WriteString(fe.err.Field() + " is required")
		return sb.String()
	}

	sb.WriteString("validation failed on field '" + fe.err.Field() + "'")
	sb.WriteString(", condition: " + fe.err.ActualTag())

	// Print condition parameters, e.g. oneof=red blue -> { red blue }
	if fe.err.Param() != "" {
		sb.WriteString(" { " + fe.err.Param() + " }")
	}

	if fe.err.Value() != nil && fe.err.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", fe.err.Value()))
	}

	return sb.String()
}

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		response := transport.BaseResponse{
			Success: false,
			Data:    nil,
			Errors:  nil,
		}

		if err != nil {
			if validationErr, ok := err.(validator.ValidationErrors); ok {
				for _, fieldErr := range validationErr {
					response.Errors = append(response.Errors, transport.ErrorBody{
						Message: FieldValidationError{fieldErr}.Error(),
					})
				}
				return c.JSON(400, response)
			} else {
				if httpErr, ok := err.(*echo.HTTPError); ok {
					response.Errors = append(response.Errors, transport.ErrorBody{
						Message: httpErr.Message.(string),
					})
					return c.JSON(httpErr.Code, response)
				}

				if errBody, ok := err.(transport.ErrorBody); ok {
					response.Errors = append(response.Errors, errBody)
					return c.JSON(errBody.HttpCode, response)
				}
				
				response.Errors = append(response.Errors, transport.ErrorBody{
					Message: err.Error(),
				})
				return c.JSON(500, response)
			}
		}
		return nil
	}

}
