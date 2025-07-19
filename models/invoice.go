package models

import "time"

type Invoice struct {
	InvoiceId              int       `json:"invoice_id" gorm:"primaryKey"`
	UserUuid               string    `json:"user_uuid" gorm:"type:varchar(200)"`
	CartCode               string    `json:"cart_code" gorm:"type:varchar(200)"`
	InvoiceNo              string    `json:"invoice_no" gorm:"type:varchar(150)"`
	InvoiceDesc            string    `json:"invoice_desc" gorm:"type:text"`
	InvoicePrice           string    `json:"invoice_price" gorm:"type:varchar(30)"`
	InvoiceStatusCd        string    `json:"invoice_status_cd" type:"varchar(30)"`
	InvoiceCreatedDate     time.Time `json:"created_invoice_date"`
	InvoiceCreatedUserUuid string    `json:"created_invoice_user_uuid" gorm:"type:varchar(200)"`
	InvoiceCreatedUsername string    `json:"created_invoice_user_username" gorm:"type:varchar(100)"`
	InvoiceUpdatedDate     time.Time `json:"updated_invoice_date"`
	InvoiceUpdatedUserUuid string    `json:"updated_invoice_user_uuid" gorm:"type:varchar(200)"`
	InvoiceUpdatedUsername string    `json:"updated_invoice_user_username" gorm:"type:varchar(100)"`
}
