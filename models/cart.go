package models

import "time"

type Cart struct {
	CartId              int       `json:"cart_id" gorm:"primaryKey"`
	UserUuid            string    `json:"user_uuid" gorm:"type:varchar(200)"`
	ProductCode         string    `json:"product_code" gorm:"type:varchar(200)"`
	CartCode            string    `json:"cart_code" gorm:"type:varchar(200)"`
	CartDescription     string    `json:"cart_description" gorm:"type:text"`
	CartPrice           string    `json:"cart_price" gorm:"type:varchar(30)"`
	CartQuantity        string    `json:"cart_quantity" gorm:"type:varchar(30)"`
	CartStatusCd        string    `json:"cart_status_cd"  gorm:"type:varchar(30)"`
	CartCreatedDate     time.Time `json:"created_cart_date"`
	CartCreatedUserUuid string    `json:"created_cart_user_uuid" gorm:"type:varchar(200)"`
	CartCreatedUsername string    `json:"created_cart_user_username" gorm:"type:varchar(100)"`
	CartUpdatedDate     time.Time `json:"updated_cart_date"`
	CartUpdatedUserUuid string    `json:"updated_cart_user_uuid" gorm:"type:varchar(200)"`
	CartUpdatedUsername string    `json:"updated_cart_user_username" gorm:"type:varchar(100)"`
}
