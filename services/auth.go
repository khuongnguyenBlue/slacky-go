package services

import (
	"crypto/rsa"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/khuongnguyenBlue/slacky/constants"
	"github.com/khuongnguyenBlue/slacky/interfaces/repositories"
	"github.com/khuongnguyenBlue/slacky/models"
	"github.com/khuongnguyenBlue/slacky/transport"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repositories.IUserRepository
	jwtPrivateKey *rsa.PrivateKey
}

type CustomClaims struct {
	UserId uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthService(userRepo repositories.IUserRepository) *AuthService {
	pemString := viper.GetString("JWT_SECRET")
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pemString))
	if err != nil {
		panic(err)
	}

	return &AuthService{userRepo, privateKey}
}

func (service *AuthService) Register(request transport.RegistrationRequest) (*models.User, error) {
	hash, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	user, err := service.CreateUser(request.Email, hash)
	if err != nil {
		if errBody, ok := err.(transport.ErrorBody); ok {
			if errBody.Code == constants.DB_SQLSTATE_UNIQUE_VIOLATION {
				errBody.HttpCode = 400
				errBody.Message = "email already exists"
				return nil, errBody
			}
		}

		return nil, err
	}

	return user, nil
}

func (service *AuthService) Login(request transport.LoginRequest) (string, error) {
	user, err := service.GetUserByEmail(request.Email)
	if err != nil {
		return "", err
	}

	if user.ID == 0 {
		return "", echo.NewHTTPError(400, "user not found")
	}

	if !checkPasswordHash(request.Password, user.Hash) {
		return "", echo.NewHTTPError(400, "wrong password")
	}

	return service.generateToken(*user)
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (service *AuthService) generateToken(user models.User) (string, error) {
	claims := CustomClaims{
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			Issuer:    "slacky",
			Subject:   user.Email,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256.SigningMethodRSA, claims)
	return token.SignedString(service.jwtPrivateKey)
}
