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

type cartController struct {
	cartService services.CartService
}

func NewCartController(cartService services.CartService) *cartController {
	return &cartController{cartService}
}

func (cc *cartController) FindCartsController(c *fiber.Ctx) error {
	carts, err := cc.cartService.FindAllService()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All carts are empty!",
		})
	}

	if len(carts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All carts are empty! Please try again!",
		})
	}

	var cartResponse []responses.CartResponse

	for _, cart := range carts {
		cartRsps := responses.GetCartResponse(cart)

		cartResponse = append(cartResponse, cartRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All carts are succesfully appeared!",
		"data":    cartResponse,
	})
}

func (cc *cartController) FindCartsByUserIdController(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)
	user_uuid := currentUser["user_uuid"]

	carts, err := cc.cartService.FindAllByUserIdService(user_uuid)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All carts by user are empty!",
		})
	}

	if len(carts) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All carts are empty! Please try again!",
		})
	}

	var cartResponse []responses.CartResponse

	for _, cart := range carts {
		cartRsps := responses.GetCartResponse(cart)

		cartResponse = append(cartResponse, cartRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All carts by user are succesfully appeared!",
		"data":    cartResponse,
	})
}

func (cc *cartController) FindCartController(c *fiber.Ctx) error {
	cartCode := c.Params("cart_code")

	cart, err := cc.cartService.FindOneService(cartCode)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A cart %s data are not found!", cartCode),
		})
	}

	cartRsps := responses.GetCartResponse(cart)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A cart %s data are succesfully appeared!", cartCode),
		"data":    cartRsps,
	})
}

func (cc *cartController) CreateCartController(c *fiber.Ctx) error {
	var createCartInput inputs.CreateCartInput

	err := c.BodyParser(&createCartInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	cartCode := helper.GenerateUUID(createCartInput.ProductCode)

	_, errCheckCart := cc.cartService.FindOneService(cartCode)

	if errCheckCart == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A cart data is exist! Please try again!",
		})
	}

	currentUser := middleware.CurrentUser(c)

	newCart, _ := cc.cartService.CreateService(createCartInput, currentUser)

	createCartRsps := responses.GetCreateCartResponse(newCart)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A new cart is succesfully created!",
		"data":    createCartRsps,
	})
}

func (cc *cartController) UpdateCartController(c *fiber.Ctx) error {
	cartCode := c.Params("cart_code")

	var updateCartInput inputs.UpdateCartInput

	err := c.BodyParser(&updateCartInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckCart := cc.cartService.FindOneService(cartCode)

	if errCheckCart != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A cart %s data are not found!", cartCode),
		})
	}

	currentUser := middleware.CurrentUser(c)

	newUpdateCart, _ := cc.cartService.UpdateService(cartCode, updateCartInput, currentUser)

	updateCartRsps := responses.GetUpdateCartResponse(newUpdateCart)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A cart %s data are succesfully updated!", cartCode),
		"data":    updateCartRsps,
	})
}

func (cc *cartController) DeleteCartController(c *fiber.Ctx) error {
	cartCode := c.Params("cart_code")

	_, err := cc.cartService.FindOneService(cartCode)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A cart %s data are not found!", cartCode),
		})
	}

	cc.cartService.DeleteService(cartCode)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A cart %s data are succesfully deleted!", cartCode),
	})
}
