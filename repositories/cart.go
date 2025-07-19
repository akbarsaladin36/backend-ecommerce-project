package repositories

import (
	"backend-restapi-ecommerce/models"

	"gorm.io/gorm"
)

type CartRepository interface {
	FindAll() ([]models.Cart, error)
	FindAllByUserId(user_uuid string) ([]models.Cart, error)
	FindOne(cart_code string) (models.Cart, error)
	FindProduct(product_code string) (models.Product, error)
	Create(cart models.Cart) (models.Cart, error)
	Update(cart models.Cart) (models.Cart, error)
	Delete(cart models.Cart) (models.Cart, error)
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *cartRepository {
	return &cartRepository{db}
}

func (cr *cartRepository) FindAll() ([]models.Cart, error) {
	var carts []models.Cart

	err := cr.db.Find(&carts).Error

	return carts, err
}

func (cr *cartRepository) FindAllByUserId(user_uuid string) ([]models.Cart, error) {
	var carts []models.Cart

	err := cr.db.Where("user_uuid = ?", user_uuid).Find(&carts).Error

	return carts, err
}

func (cr *cartRepository) FindOne(cart_code string) (models.Cart, error) {
	var cart models.Cart

	err := cr.db.Where("cart_code = ?", cart_code).First(&cart).Error

	return cart, err
}

func (cr *cartRepository) FindProduct(product_code string) (models.Product, error) {
	var product models.Product

	err := cr.db.Where("product_code = ?", product_code).First(&product).Error

	return product, err
}

func (cr *cartRepository) Create(cart models.Cart) (models.Cart, error) {
	err := cr.db.Create(&cart).Error

	return cart, err
}

func (cr *cartRepository) Update(cart models.Cart) (models.Cart, error) {
	err := cr.db.Save(&cart).Error

	return cart, err
}

func (cr *cartRepository) Delete(cart models.Cart) (models.Cart, error) {
	err := cr.db.Delete(&cart).Error

	return cart, err
}
