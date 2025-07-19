package services

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/models"
	"backend-restapi-ecommerce/repositories"
	"time"
)

type ProductService interface {
	FindAllService() ([]models.Product, error)
	FindOneService(product_code string) (models.Product, error)
	CreateService(createProductInput inputs.CreateProductInput, currentUser map[string]string) (models.Product, error)
	UpdateService(product_code string, updateProductInput inputs.UpdateProductInput, currentUser map[string]string) (models.Product, error)
	DeleteService(product_code string) (models.Product, error)
}

type productService struct {
	productRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) *productService {
	return &productService{productRepository}
}

func (ps *productService) FindAllService() ([]models.Product, error) {
	products, err := ps.productRepository.FindAll()

	return products, err
}

func (ps *productService) FindOneService(product_code string) (models.Product, error) {
	product, err := ps.productRepository.FindOne(product_code)

	return product, err
}

func (ps *productService) CreateService(createProductInput inputs.CreateProductInput, currentUser map[string]string) (models.Product, error) {
	productCode := helper.GenerateSlug(createProductInput.ProductName)

	product := models.Product{
		ProductCode:            productCode,
		ProductName:            createProductInput.ProductName,
		ProductDescription:     createProductInput.ProductDescription,
		ProductPrice:           createProductInput.ProductPrice,
		ProductQuantity:        createProductInput.ProductQuantity,
		ProductStatusCd:        "active",
		ProductCreatedDate:     time.Now(),
		ProductCreatedUserUuid: currentUser["user_uuid"],
		ProductCreatedUsername: currentUser["user_username"],
	}

	newProduct, err := ps.productRepository.Create(product)

	return newProduct, err
}

func (ps *productService) UpdateService(product_code string, updateProductInput inputs.UpdateProductInput, currentUser map[string]string) (models.Product, error) {
	checkProduct, _ := ps.productRepository.FindOne(product_code)

	updateProductCode := helper.GenerateSlug(updateProductInput.ProductName)

	checkProduct.ProductCode = updateProductCode
	checkProduct.ProductName = updateProductInput.ProductName
	checkProduct.ProductDescription = updateProductInput.ProductDescription
	checkProduct.ProductPrice = updateProductInput.ProductPrice
	checkProduct.ProductQuantity = updateProductInput.ProductQuantity
	checkProduct.ProductUpdatedDate = time.Now()
	checkProduct.ProductUpdatedUserUuid = currentUser["user_uuid"]
	checkProduct.ProductUpdatedUsername = currentUser["user_username"]

	updateProduct, err := ps.productRepository.Update(checkProduct)

	return updateProduct, err
}

func (ps *productService) DeleteService(product_code string) (models.Product, error) {
	checkProduct, _ := ps.productRepository.FindOne(product_code)

	deleteProduct, err := ps.productRepository.Delete(checkProduct)

	return deleteProduct, err
}
