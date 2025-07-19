package services

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/models"
	"backend-restapi-ecommerce/repositories"
	"fmt"
	"strconv"
	"time"
)

type InvoiceService interface {
	FindAllService() ([]models.Invoice, error)
	FindAllByUserIdService(user_uuid string) ([]models.Invoice, error)
	FindOneService(invoice_no string) (models.Invoice, error)
	FindCartService(cart_code string) (models.Cart, error)
	FindUserService(user_uuid string) (models.User, error)
	CreateService(createInvoiceInput inputs.CreateInvoiceInput, currentUser map[string]string) (models.Invoice, error)
	UpdateService(invoice_no string, updateInvoiceInput inputs.UpdateInvoiceInput, currentUser map[string]string) (models.Invoice, error)
	DeleteService(invoice_no string) (models.Invoice, error)
}

type invoiceService struct {
	invoiceRepository repositories.InvoiceRepository
}

func NewInvoiceService(invoiceRepository repositories.InvoiceRepository) *invoiceService {
	return &invoiceService{invoiceRepository}
}

func (is *invoiceService) FindAllService() ([]models.Invoice, error) {
	invoices, err := is.invoiceRepository.FindAll()

	return invoices, err
}

func (is *invoiceService) FindAllByUserIdService(user_uuid string) ([]models.Invoice, error) {
	invoices, err := is.invoiceRepository.FindByUserId(user_uuid)

	return invoices, err
}

func (is *invoiceService) FindOneService(invoice_no string) (models.Invoice, error) {
	invoice, err := is.invoiceRepository.FindOne(invoice_no)

	return invoice, err
}

func (is *invoiceService) FindCartService(cart_code string) (models.Cart, error) {
	cart, err := is.invoiceRepository.FindCart(cart_code)

	return cart, err
}

func (is *invoiceService) FindUserService(user_uuid string) (models.User, error) {
	user, err := is.invoiceRepository.FindUser(user_uuid)

	return user, err
}

func (is *invoiceService) CreateService(createInvoiceInput inputs.CreateInvoiceInput, currentUser map[string]string) (models.Invoice, error) {
	invoiceNo := helper.GenerateInvoice()

	invoice := models.Invoice{
		UserUuid:               currentUser["user_uuid"],
		CartCode:               createInvoiceInput.CartCode,
		InvoiceNo:              invoiceNo,
		InvoiceDesc:            createInvoiceInput.InvoiceDesc,
		InvoicePrice:           createInvoiceInput.InvoicePrice,
		InvoiceStatusCd:        "pending",
		InvoiceCreatedDate:     time.Now(),
		InvoiceCreatedUserUuid: currentUser["user_uuid"],
		InvoiceCreatedUsername: currentUser["user_username"],
	}

	createInvoice, err := is.invoiceRepository.Create(invoice)

	return createInvoice, err
}

func (is *invoiceService) UpdateService(invoice_no string, updateInvoiceInput inputs.UpdateInvoiceInput, currentUser map[string]string) (models.Invoice, error) {
	checkInvoice, _ := is.invoiceRepository.FindOne(invoice_no)

	checkCart, _ := is.invoiceRepository.FindCart(checkInvoice.CartCode)

	checkUser, _ := is.invoiceRepository.FindUser(checkInvoice.UserUuid)

	checkProduct, _ := is.invoiceRepository.FindProduct(checkCart.ProductCode)

	checkInvoice.InvoiceStatusCd = updateInvoiceInput.InvoiceStatusCd
	checkInvoice.InvoiceUpdatedDate = time.Now()
	checkInvoice.InvoiceUpdatedUserUuid = currentUser["user_uuid"]
	checkInvoice.InvoiceUpdatedUsername = currentUser["user_username"]

	updateInvoice, err := is.invoiceRepository.Update(checkInvoice)

	switch updateInvoiceInput.InvoiceStatusCd {
	case "paying":
		checkCart.CartStatusCd = updateInvoiceInput.InvoiceStatusCd
		checkCart.CartUpdatedDate = time.Now()
		checkCart.CartUpdatedUserUuid = currentUser["user_uuid"]
		checkCart.CartUpdatedUsername = currentUser["user_username"]

		is.invoiceRepository.UpdateCart(checkCart)
	case "paid":
		checkCart.CartStatusCd = updateInvoiceInput.InvoiceStatusCd
		checkCart.CartUpdatedDate = time.Now()
		checkCart.CartUpdatedUserUuid = currentUser["user_uuid"]
		checkCart.CartUpdatedUsername = currentUser["user_username"]

		newUserBalanceTransaction, _ := strconv.ParseInt(checkUser.UserBalanceTransactionAmount, 0, 64)
		checkInvoicePrice, _ := strconv.ParseInt(checkInvoice.InvoicePrice, 0, 64)
		calculateBalanceTransaction := newUserBalanceTransaction - checkInvoicePrice
		formattedUserBalanceTransaction := strconv.FormatInt(calculateBalanceTransaction, 10)

		cartQuantity, _ := strconv.ParseInt(checkCart.CartQuantity, 0, 64)
		productQuantity, _ := strconv.ParseInt(checkProduct.ProductQuantity, 0, 64)
		calculateQuantity := productQuantity - cartQuantity
		formattedQuantity := strconv.FormatInt(calculateQuantity, 10)

		checkUser.UserBalanceTransactionAmount = formattedUserBalanceTransaction

		checkProduct.ProductQuantity = formattedQuantity

		is.invoiceRepository.UpdateCart(checkCart)

		is.invoiceRepository.UpdateProduct(checkProduct)

		is.invoiceRepository.UpdateUserBalanceTransaction(checkUser)
	default:
		fmt.Println("Invalid Process")
	}

	return updateInvoice, err
}

func (is *invoiceService) DeleteService(invoice_no string) (models.Invoice, error) {
	checkInvoice, _ := is.invoiceRepository.FindOne(invoice_no)

	deleteInvoice, err := is.invoiceRepository.Delete(checkInvoice)

	return deleteInvoice, err
}
