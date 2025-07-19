package controllers

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/middleware"
	"backend-restapi-ecommerce/responses"
	"backend-restapi-ecommerce/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type productController struct {
	productService services.ProductService
}

func NewProductController(productService services.ProductService) *productController {
	return &productController{productService}
}

func (pc *productController) FindProductsController(c *fiber.Ctx) error {
	products, err := pc.productService.FindAllService()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All products are empty!",
		})
	}

	if len(products) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All products are empty!",
		})
	}

	var productResponse []responses.ProductResponse

	for _, product := range products {
		productRsps := responses.GetProductResponse(product)

		productResponse = append(productResponse, productRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All products are succesfully appeared!",
		"data":    productResponse,
	})
}

func (pc *productController) FindProductController(c *fiber.Ctx) error {
	product_code := c.Params("product_code")

	product, err := pc.productService.FindOneService(product_code)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A product %s data are not found!", product_code),
		})
	}

	productRsps := responses.GetProductResponse(product)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A product %s data are succesfully appeared!", product_code),
		"data":    productRsps,
	})
}

func (pc *productController) CreateProductController(c *fiber.Ctx) error {
	var createProductInput inputs.CreateProductInput

	err := c.BodyParser(&createProductInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	productCode := helper.GenerateSlug(createProductInput.ProductName)

	_, errCheckProduct := pc.productService.FindOneService(productCode)

	if errCheckProduct == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A product has been registered! Please try again!",
		})
	}

	currentUser := middleware.CurrentUser(c)

	newProduct, _ := pc.productService.CreateService(createProductInput, currentUser)

	createProductRsps := responses.GetCreateProductResponse(newProduct)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A new product is succesfully created!",
		"data":    createProductRsps,
	})
}

func (pc *productController) UpdateProductController(c *fiber.Ctx) error {
	product_code := c.Params("product_code")

	var updateProductInput inputs.UpdateProductInput

	err := c.BodyParser(&updateProductInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckProduct := pc.productService.FindOneService(product_code)

	if errCheckProduct != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A product has not been registered! Please try again!",
		})
	}

	currentUser := middleware.CurrentUser(c)

	updateProduct, _ := pc.productService.UpdateService(product_code, updateProductInput, currentUser)

	updateProductRsps := responses.GetUpdateProductResponse(updateProduct)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A product %s data are succesfully updated!", product_code),
		"data":    updateProductRsps,
	})
}

func (pc *productController) DeleteProductController(c *fiber.Ctx) error {
	product_code := c.Params("product_code")

	_, err := pc.productService.FindOneService(product_code)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A product %s data are not found!", product_code),
		})
	}

	pc.productService.DeleteService(product_code)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A product %s data are succesfully deleted!", product_code),
	})

}
