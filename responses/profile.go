package responses

import (
	"backend-restapi-ecommerce/models"
)

type ProfileResponse struct {
	UserFirstName                string `json:"user_first_name"`
	UserLastName                 string `json:"user_last_name"`
	UserAddress                  string `json:"user_address"`
	UserPhoneNumber              string `json:"user_phone_number"`
	UserBalanceTransactionAmount string `json:"user_balance_transaction_amount"`
}

type ProfileBalanceTransactionResponse struct {
	ProfileResponse
	UserBalanceTransaction string `json:"user_balance_transaction"`
}

func GetProfileResponse(userRsps models.User) ProfileResponse {
	return ProfileResponse{
		UserFirstName:                userRsps.UserFirstName,
		UserLastName:                 userRsps.UserLastName,
		UserAddress:                  userRsps.UserAddress,
		UserPhoneNumber:              userRsps.UserPhoneNumber,
		UserBalanceTransactionAmount: userRsps.UserBalanceTransactionAmount,
	}
}

func GetProfileBalanceTransactionResponse(userRsps models.User) ProfileBalanceTransactionResponse {
	return ProfileBalanceTransactionResponse{
		ProfileResponse: ProfileResponse{
			UserFirstName:   userRsps.UserFirstName,
			UserLastName:    userRsps.UserLastName,
			UserAddress:     userRsps.UserAddress,
			UserPhoneNumber: userRsps.UserPhoneNumber,
		},
		UserBalanceTransaction: userRsps.UserBalanceTransactionAmount,
	}
}
