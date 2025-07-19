package responses

import (
	"backend-restapi-ecommerce/models"
	"time"
)

type ProductResponse struct {
	ProductId          int    `json:"product_id"`
	ProductCode        string `json:"product_code"`
	ProductName        string `json:"product_name"`
	ProductDescription string `json:"product_description"`
	ProductPrice       string `json:"product_price"`
	ProductQuantity    string `json:"product_quantity"`
	ProductStatusCd    string `json:"product_status_cd"`
}

type CreateProductResponse struct {
	ProductCode            string    `json:"product_code"`
	ProductName            string    `json:"product_name"`
	ProductDescription     string    `json:"product_description"`
	ProductPrice           string    `json:"product_price"`
	ProductQuantity        string    `json:"product_quantity"`
	ProductStatusCd        string    `json:"product_status_cd"`
	ProductCreatedDate     time.Time `json:"created_product_date"`
	ProductCreatedUserUuid string    `json:"created_product_user_uuid"`
	ProductCreatedUsername string    `json:"created_product_user_username"`
}

type UpdateProductResponse struct {
	ProductCode            string    `json:"product_code"`
	ProductName            string    `json:"product_name"`
	ProductDescription     string    `json:"product_description"`
	ProductPrice           string    `json:"product_price"`
	ProductQuantity        string    `json:"product_quantity"`
	ProductStatusCd        string    `json:"product_status_cd"`
	ProductUpdatedDate     time.Time `json:"updated_product_date"`
	ProductUpdatedUserUuid string    `json:"updated_product_user_uuid"`
	ProductUpdatedUsername string    `json:"updated_product_user_username"`
}

func GetProductResponse(productRsps models.Product) ProductResponse {
	return ProductResponse{
		ProductId:          productRsps.ProductId,
		ProductCode:        productRsps.ProductCode,
		ProductName:        productRsps.ProductName,
		ProductDescription: productRsps.ProductDescription,
		ProductPrice:       productRsps.ProductPrice,
		ProductQuantity:    productRsps.ProductQuantity,
		ProductStatusCd:    productRsps.ProductStatusCd,
	}
}

func GetCreateProductResponse(productRsps models.Product) CreateProductResponse {
	return CreateProductResponse{
		ProductCode:            productRsps.ProductCode,
		ProductName:            productRsps.ProductName,
		ProductDescription:     productRsps.ProductDescription,
		ProductPrice:           productRsps.ProductPrice,
		ProductQuantity:        productRsps.ProductQuantity,
		ProductStatusCd:        productRsps.ProductStatusCd,
		ProductCreatedDate:     productRsps.ProductCreatedDate,
		ProductCreatedUserUuid: productRsps.ProductCreatedUserUuid,
		ProductCreatedUsername: productRsps.ProductCreatedUsername,
	}
}

func GetUpdateProductResponse(productRsps models.Product) UpdateProductResponse {
	return UpdateProductResponse{
		ProductCode:            productRsps.ProductCode,
		ProductName:            productRsps.ProductName,
		ProductDescription:     productRsps.ProductDescription,
		ProductPrice:           productRsps.ProductPrice,
		ProductQuantity:        productRsps.ProductQuantity,
		ProductStatusCd:        productRsps.ProductStatusCd,
		ProductUpdatedDate:     productRsps.ProductUpdatedDate,
		ProductUpdatedUserUuid: productRsps.ProductUpdatedUserUuid,
		ProductUpdatedUsername: productRsps.ProductUpdatedUsername,
	}
}
