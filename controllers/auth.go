package controllers

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/responses"
	"backend-restapi-ecommerce/services"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *authController {
	return &authController{authService}
}

func (ac *authController) RegisterController(c *fiber.Ctx) error {
	var registerInput inputs.RegisterInput

	err := c.BodyParser(&registerInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckUser := ac.authService.FindOneService(registerInput.Username)

	if errCheckUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A username has been registered! Please try again!",
		})
	}

	newRegisterUser, _ := ac.authService.RegisterService(registerInput)

	registerUserRsps := responses.GetRegisterResponse(newRegisterUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A new user is succesfully created!",
		"data":    registerUserRsps,
	})
}

func (ac *authController) LoginController(c *fiber.Ctx) error {
	var loginInput inputs.LoginInput

	err := c.BodyParser(&loginInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckUser := ac.authService.FindOneService(loginInput.Username)

	if errCheckUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A username has not been registered! Please register a user now!",
		})
	}

	loginUser, errLoginUser := ac.authService.LoginService(loginInput)

	if errLoginUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A password is not match! Please try again!",
		})
	}

	tokenString, _ := helper.GenerateToken(loginUser)

	loginUserRsps := responses.GetLoginResponse(loginUser, tokenString)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A user is succesfully login!",
		"data":    loginUserRsps,
	})
}
