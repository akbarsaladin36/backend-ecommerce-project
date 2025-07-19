package responses

import (
	"backend-restapi-ecommerce/models"
	"time"
)

type InvoiceResponse struct {
	InvoiceId          int    `json:"invoice_id"`
	UserUuid           string `json:"user_uuid"`
	CartCode           string `json:"cart_code"`
	InvoiceNo          string `json:"invoice_no"`
	InvoiceDescription string `json:"invoice_description"`
	InvoicePrice       string `json:"invoice_price"`
	InvoiceStatusCd    string `json:"invoice_status_cd"`
}

type CreateInvoiceResponse struct {
	UserUuid               string    `json:"user_uuid"`
	CartCode               string    `json:"cart_code"`
	InvoiceNo              string    `json:"invoice_no"`
	InvoiceDescription     string    `json:"invoice_description"`
	InvoicePrice           string    `json:"invoice_price"`
	InvoiceStatusCd        string    `json:"invoice_status_cd"`
	InvoiceCreatedDate     time.Time `json:"created_invoice_date"`
	InvoiceCreatedUserUuid string    `json:"created_invoice_user_uuid"`
	InvoiceCreatedUsername string    `json:"created_invoice_user_username"`
}

type UpdateInvoiceResponse struct {
	UserUuid               string    `json:"user_uuid"`
	CartCode               string    `json:"cart_code"`
	InvoiceNo              string    `json:"invoice_no"`
	InvoiceDescription     string    `json:"invoice_description"`
	InvoicePrice           string    `json:"invoice_price"`
	InvoiceStatusCd        string    `json:"invoice_status_cd"`
	InvoiceUpdatedDate     time.Time `json:"updated_invoice_date"`
	InvoiceUpdatedUserUuid string    `json:"updated_invoice_user_uuid"`
	InvoiceUpdatedUsername string    `json:"updated_invoice_user_username"`
}

func GetInvoiceResponse(invoiceRsps models.Invoice) InvoiceResponse {
	return InvoiceResponse{
		InvoiceId:          invoiceRsps.InvoiceId,
		UserUuid:           invoiceRsps.UserUuid,
		CartCode:           invoiceRsps.CartCode,
		InvoiceNo:          invoiceRsps.InvoiceNo,
		InvoiceDescription: invoiceRsps.InvoiceDesc,
		InvoicePrice:       invoiceRsps.InvoicePrice,
		InvoiceStatusCd:    invoiceRsps.InvoiceStatusCd,
	}
}

func GetCreateInvoiceResponse(invoiceRsps models.Invoice) CreateInvoiceResponse {
	return CreateInvoiceResponse{
		UserUuid:               invoiceRsps.UserUuid,
		CartCode:               invoiceRsps.CartCode,
		InvoiceNo:              invoiceRsps.InvoiceNo,
		InvoiceDescription:     invoiceRsps.InvoiceDesc,
		InvoicePrice:           invoiceRsps.InvoicePrice,
		InvoiceStatusCd:        invoiceRsps.InvoiceStatusCd,
		InvoiceCreatedDate:     invoiceRsps.InvoiceCreatedDate,
		InvoiceCreatedUserUuid: invoiceRsps.InvoiceCreatedUserUuid,
		InvoiceCreatedUsername: invoiceRsps.InvoiceCreatedUsername,
	}
}

func GetUpdateInvoiceResponse(invoiceRsps models.Invoice) UpdateInvoiceResponse {
	return UpdateInvoiceResponse{
		UserUuid:               invoiceRsps.UserUuid,
		CartCode:               invoiceRsps.CartCode,
		InvoiceNo:              invoiceRsps.InvoiceNo,
		InvoiceDescription:     invoiceRsps.InvoiceDesc,
		InvoicePrice:           invoiceRsps.InvoicePrice,
		InvoiceStatusCd:        invoiceRsps.InvoiceStatusCd,
		InvoiceUpdatedDate:     invoiceRsps.InvoiceUpdatedDate,
		InvoiceUpdatedUserUuid: invoiceRsps.InvoiceUpdatedUserUuid,
		InvoiceUpdatedUsername: invoiceRsps.InvoiceUpdatedUsername,
	}
}
