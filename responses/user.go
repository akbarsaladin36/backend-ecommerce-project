package responses

import (
	"backend-restapi-ecommerce/models"
	"time"
)

type UserResponse struct {
	UserUuid        string `json:"user_uuid"`
	UserUsername    string `json:"user_username"`
	UserEmail       string `json:"user_email"`
	UserFirstName   string `json:"user_first_name"`
	UserLastName    string `json:"user_last_name"`
	UserAddress     string `json:"user_address"`
	UserPhoneNumber string `json:"user_phone_number"`
	UserRole        string `json:"user_role"`
	UserStatusCd    string `json:"user_status_cd"`
}

type CreateUserResponse struct {
	UserUuid            string    `json:"user_uuid"`
	UserUsername        string    `json:"user_username"`
	UserEmail           string    `json:"user_email"`
	UserFirstName       string    `json:"user_first_name"`
	UserLastName        string    `json:"user_last_name"`
	UserAddress         string    `json:"user_address"`
	UserPhoneNumber     string    `json:"user_phone_number"`
	UserCreatedDate     time.Time `json:"created_user_date"`
	UserCreatedUserUuid string    `json:"created_user_uuid"`
	UserCreatedUsername string    `json:"created_user_username"`
}

type UpdateUserResponse struct {
	UserFirstName       string    `json:"user_first_name"`
	UserLastName        string    `json:"user_last_name"`
	UserAddress         string    `json:"user_address"`
	UserPhoneNumber     string    `json:"user_phone_number"`
	UserUpdatedDate     time.Time `json:"updated_user_date"`
	UserUpdatedUserUuid string    `json:"updated_user_uuid"`
	UserUpdatedUsername string    `json:"updated_user_username"`
}
type UpdateUserStatusResponse struct {
	UserStatusCd        string    `json:"user_status_cd"`
	UserUpdatedDate     time.Time `json:"updated_user_date"`
	UserUpdatedUserUuid string    `json:"updated_user_uuid"`
	UserUpdatedUsername string    `json:"updated_user_username"`
}

func GetUserResponse(userRsps models.User) UserResponse {
	return UserResponse{
		UserUuid:        userRsps.UserUUID,
		UserUsername:    userRsps.UserUsername,
		UserEmail:       userRsps.UserEmail,
		UserFirstName:   userRsps.UserFirstName,
		UserLastName:    userRsps.UserLastName,
		UserAddress:     userRsps.UserAddress,
		UserPhoneNumber: userRsps.UserPhoneNumber,
		UserRole:        userRsps.UserRole,
		UserStatusCd:    userRsps.UserStatusCd,
	}
}

func GetCreateUserResponse(userRsps models.User) CreateUserResponse {
	return CreateUserResponse{
		UserUuid:            userRsps.UserUUID,
		UserUsername:        userRsps.UserUsername,
		UserEmail:           userRsps.UserEmail,
		UserFirstName:       userRsps.UserFirstName,
		UserLastName:        userRsps.UserLastName,
		UserAddress:         userRsps.UserAddress,
		UserPhoneNumber:     userRsps.UserPhoneNumber,
		UserCreatedDate:     userRsps.UserCreatedDate,
		UserCreatedUserUuid: userRsps.UserCreatedUserUuid,
		UserCreatedUsername: userRsps.UserCreatedUsername,
	}
}

func GetUpdateUserResponse(userRsps models.User) UpdateUserResponse {
	return UpdateUserResponse{
		UserFirstName:       userRsps.UserFirstName,
		UserLastName:        userRsps.UserLastName,
		UserAddress:         userRsps.UserAddress,
		UserPhoneNumber:     userRsps.UserPhoneNumber,
		UserUpdatedDate:     userRsps.UserUpdatedDate,
		UserUpdatedUserUuid: userRsps.UserUpdatedUserUuid,
		UserUpdatedUsername: userRsps.UserUpdatedUsername,
	}
}

func GetUpdateUserStatusResponse(userRsps models.User) UpdateUserStatusResponse {
	return UpdateUserStatusResponse{
		UserStatusCd:        userRsps.UserStatusCd,
		UserUpdatedDate:     userRsps.UserUpdatedDate,
		UserUpdatedUserUuid: userRsps.UserUpdatedUserUuid,
		UserUpdatedUsername: userRsps.UserUpdatedUsername,
	}
}
