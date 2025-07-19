package responses

import (
	"backend-restapi-ecommerce/models"
	"time"
)

type CartResponse struct {
	UserUuid        string `json:"user_uuid"`
	ProductCode     string `json:"product_code"`
	CartCode        string `json:"cart_code"`
	CartDescription string `json:"cart_description"`
	CartPrice       string `json:"cart_price"`
	CartQuantity    string `json:"cart_quantity"`
	CartStatusCd    string `json:"cart_status_cd"`
}

type CreateCartResponse struct {
	UserUuid            string    `json:"user_uuid"`
	ProductCode         string    `json:"product_code"`
	CartCode            string    `json:"cart_code"`
	CartDescription     string    `json:"cart_description"`
	CartPrice           string    `json:"cart_price"`
	CartQuantity        string    `json:"cart_quantity"`
	CartStatusCd        string    `json:"cart_status_cd"`
	CartCreatedDate     time.Time `json:"created_cart_date"`
	CartCreatedUserUuid string    `json:"created_cart_user_uuid"`
	CartCreatedUsername string    `json:"created_cart_user_username"`
}

type UpdateCartResponse struct {
	UserUuid            string    `json:"user_uuid"`
	ProductCode         string    `json:"product_code"`
	CartCode            string    `json:"cart_code"`
	CartDescription     string    `json:"cart_description"`
	CartPrice           string    `json:"cart_price"`
	CartQuantity        string    `json:"cart_quantity"`
	CartStatusCd        string    `json:"cart_status_cd"`
	CartUpdatedDate     time.Time `json:"updated_cart_date"`
	CartUpdatedUserUuid string    `json:"updated_cart_user_uuid"`
	CartUpdatedUsername string    `json:"updated_cart_user_username"`
}

func GetCartResponse(cartRsps models.Cart) CartResponse {
	return CartResponse{
		UserUuid:        cartRsps.UserUuid,
		ProductCode:     cartRsps.ProductCode,
		CartCode:        cartRsps.CartCode,
		CartDescription: cartRsps.CartDescription,
		CartPrice:       cartRsps.CartPrice,
		CartQuantity:    cartRsps.CartQuantity,
		CartStatusCd:    cartRsps.CartStatusCd,
	}
}

func GetCreateCartResponse(cartRsps models.Cart) CreateCartResponse {
	return CreateCartResponse{
		UserUuid:            cartRsps.UserUuid,
		ProductCode:         cartRsps.ProductCode,
		CartCode:            cartRsps.CartCode,
		CartDescription:     cartRsps.CartDescription,
		CartPrice:           cartRsps.CartPrice,
		CartQuantity:        cartRsps.CartQuantity,
		CartStatusCd:        cartRsps.CartStatusCd,
		CartCreatedDate:     cartRsps.CartCreatedDate,
		CartCreatedUserUuid: cartRsps.CartCreatedUserUuid,
		CartCreatedUsername: cartRsps.CartCreatedUsername,
	}
}

func GetUpdateCartResponse(cartRsps models.Cart) UpdateCartResponse {
	return UpdateCartResponse{
		UserUuid:            cartRsps.UserUuid,
		ProductCode:         cartRsps.ProductCode,
		CartCode:            cartRsps.CartCode,
		CartDescription:     cartRsps.CartDescription,
		CartPrice:           cartRsps.CartPrice,
		CartQuantity:        cartRsps.CartQuantity,
		CartStatusCd:        cartRsps.CartStatusCd,
		CartUpdatedDate:     cartRsps.CartUpdatedDate,
		CartUpdatedUserUuid: cartRsps.CartUpdatedUserUuid,
		CartUpdatedUsername: cartRsps.CartUpdatedUsername,
	}
}
