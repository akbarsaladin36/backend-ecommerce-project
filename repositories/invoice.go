package repositories

import (
	"backend-restapi-ecommerce/models"

	"gorm.io/gorm"
)

type InvoiceRepository interface {
	FindAll() ([]models.Invoice, error)
	FindByUserId(user_uuid string) ([]models.Invoice, error)
	FindOne(invoice_no string) (models.Invoice, error)
	FindCart(cart_code string) (models.Cart, error)
	FindUser(user_uuid string) (models.User, error)
	FindProduct(product_code string) (models.Product, error)
	Create(invoice models.Invoice) (models.Invoice, error)
	Update(invoice models.Invoice) (models.Invoice, error)
	UpdateCart(cart models.Cart) error
	UpdateUserBalanceTransaction(user models.User) error
	UpdateProduct(product models.Product) error
	Delete(invoice models.Invoice) (models.Invoice, error)
}

type invoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *invoiceRepository {
	return &invoiceRepository{db}
}

func (ir *invoiceRepository) FindAll() ([]models.Invoice, error) {
	var invoices []models.Invoice

	err := ir.db.Find(&invoices).Error

	return invoices, err
}

func (ir *invoiceRepository) FindByUserId(user_uuid string) ([]models.Invoice, error) {
	var invoices []models.Invoice

	err := ir.db.Where("user_uuid = ?", user_uuid).Find(&invoices).Error

	return invoices, err
}

func (ir *invoiceRepository) FindOne(invoice_no string) (models.Invoice, error) {
	var invoice models.Invoice

	err := ir.db.Where("invoice_no = ?", invoice_no).First(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) FindCart(cart_code string) (models.Cart, error) {
	var cart models.Cart

	err := ir.db.Where("cart_code = ?", cart_code).First(&cart).Error

	return cart, err
}

func (ir *invoiceRepository) FindUser(user_uuid string) (models.User, error) {
	var user models.User

	err := ir.db.Where("user_uuid = ?", user_uuid).First(&user).Error

	return user, err
}

func (ir *invoiceRepository) FindProduct(product_code string) (models.Product, error) {
	var product models.Product

	err := ir.db.Where("product_code = ?", product_code).First(&product).Error

	return product, err
}

func (ir *invoiceRepository) Create(invoice models.Invoice) (models.Invoice, error) {
	err := ir.db.Create(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) Update(invoice models.Invoice) (models.Invoice, error) {
	err := ir.db.Save(&invoice).Error

	return invoice, err
}

func (ir *invoiceRepository) UpdateCart(cart models.Cart) error {
	err := ir.db.Save(&cart).Error

	return err
}

func (ir *invoiceRepository) UpdateUserBalanceTransaction(user models.User) error {
	err := ir.db.Save(&user).Error

	return err
}

func (ir *invoiceRepository) UpdateProduct(product models.Product) error {
	err := ir.db.Save(&product).Error

	return err
}

func (ir *invoiceRepository) Delete(invoice models.Invoice) (models.Invoice, error) {
	err := ir.db.Delete(&invoice).Error

	return invoice, err
}
