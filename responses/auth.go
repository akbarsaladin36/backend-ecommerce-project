package responses

import (
	"backend-restapi-ecommerce/models"
	"time"
)

type RegisterResponse struct {
	UserUuid            string    `json:"user_uuid"`
	UserUsername        string    `json:"user_username"`
	UserEmail           string    `json:"user_email"`
	UserCreatedDate     time.Time `json:"created_user_date"`
	UserCreatedUserUuid string    `json:"created_user_uuid"`
	UserCreatedUsername string    `json:"created_user_username"`
}

type LoginResponse struct {
	UserUuid     string `json:"user_uuid"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserToken    string `json:"user_token"`
}

func GetRegisterResponse(userRsps models.User) RegisterResponse {
	return RegisterResponse{
		UserUuid:            userRsps.UserUUID,
		UserUsername:        userRsps.UserUsername,
		UserEmail:           userRsps.UserEmail,
		UserCreatedDate:     userRsps.UserCreatedDate,
		UserCreatedUserUuid: userRsps.UserCreatedUserUuid,
		UserCreatedUsername: userRsps.UserCreatedUsername,
	}
}

func GetLoginResponse(userRsps models.User, tokenString string) LoginResponse {
	return LoginResponse{
		UserUuid:     userRsps.UserUUID,
		UserUsername: userRsps.UserUsername,
		UserEmail:    userRsps.UserEmail,
		UserToken:    tokenString,
	}
}
