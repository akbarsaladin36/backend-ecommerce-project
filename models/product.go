package models

import "time"

type Product struct {
	ProductId              int       `json:"product_id" gorm:"primaryKey"`
	ProductCode            string    `json:"product_code" gorm:"type:varchar(200)"`
	ProductName            string    `json:"product_name" gorm:"type:varchar(200)"`
	ProductDescription     string    `json:"product_description" type:"text"`
	ProductPrice           string    `json:"product_price" type:"varchar(30)"`
	ProductQuantity        string    `json:"product_quantity" type:"varchar(30)"`
	ProductStatusCd        string    `json:"product_status_cd" type:"varchar(30)"`
	ProductCreatedDate     time.Time `json:"created_product_date"`
	ProductCreatedUserUuid string    `json:"created_product_user_uuid" gorm:"type:varchar(200)"`
	ProductCreatedUsername string    `json:"created_product_user_username" gorm:"type:varchar(100)"`
	ProductUpdatedDate     time.Time `json:"updated_product_date"`
	ProductUpdatedUserUuid string    `json:"updated_product_user_uuid" gorm:"type:varchar(200)"`
	ProductUpdatedUsername string    `json:"updated_product_user_username" gorm:"type:varchar(100)"`
}
