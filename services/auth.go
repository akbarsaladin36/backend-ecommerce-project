package services

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/models"
	"backend-restapi-ecommerce/repositories"
	"time"
)

type AuthService interface {
	FindOneService(username string) (models.User, error)
	RegisterService(registerInput inputs.RegisterInput) (models.User, error)
	LoginService(loginInput inputs.LoginInput) (models.User, error)
}

type authService struct {
	authRepository repositories.AuthRepository
}

func NewAuthService(authRepository repositories.AuthRepository) *authService {
	return &authService{authRepository}
}

func (as *authService) FindOneService(username string) (models.User, error) {
	user, err := as.authRepository.FindOne(username)

	return user, err
}

func (as *authService) RegisterService(registerInput inputs.RegisterInput) (models.User, error) {
	hashedPassword, _ := helper.HashPassword(registerInput.Password)
	userUUID := helper.GenerateUUID(registerInput.Username)

	user := models.User{
		UserUUID:            userUUID,
		UserUsername:        registerInput.Username,
		UserEmail:           registerInput.Email,
		UserPassword:        hashedPassword,
		UserRole:            "user",
		UserStatusCd:        "active",
		UserCreatedDate:     time.Now(),
		UserCreatedUserUuid: userUUID,
		UserCreatedUsername: registerInput.Username,
	}

	newUser, err := as.authRepository.Create(user)

	return newUser, err
}

func (as *authService) LoginService(loginInput inputs.LoginInput) (models.User, error) {
	checkUser, err := as.authRepository.FindOne(loginInput.Username)

	helper.CheckPassword(checkUser.UserPassword, loginInput.Password)

	return checkUser, err
}
