package repositories

import (
	"backend-restapi-ecommerce/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]models.Product, error)
	FindOne(product_code string) (models.Product, error)
	Create(product models.Product) (models.Product, error)
	Update(product models.Product) (models.Product, error)
	Delete(product models.Product) (models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (pr *productRepository) FindAll() ([]models.Product, error) {
	var products []models.Product

	err := pr.db.Find(&products).Error

	return products, err
}

func (pr *productRepository) FindOne(product_code string) (models.Product, error) {
	var product models.Product

	err := pr.db.Where("product_code = ?", product_code).First(&product).Error

	return product, err
}

func (pr *productRepository) Create(product models.Product) (models.Product, error) {
	err := pr.db.Create(&product).Error

	return product, err
}

func (pr *productRepository) Update(product models.Product) (models.Product, error) {
	err := pr.db.Save(&product).Error

	return product, err
}

func (pr *productRepository) Delete(product models.Product) (models.Product, error) {
	err := pr.db.Delete(&product).Error

	return product, err
}
