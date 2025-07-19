package controllers

import (
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/middleware"
	"backend-restapi-ecommerce/responses"
	"backend-restapi-ecommerce/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type profileController struct {
	profileService services.ProfileService
}

func NewProfileController(profileService services.ProfileService) *profileController {
	return &profileController{profileService}
}

func (pc *profileController) FindProfileController(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)

	profile, err := pc.profileService.FindOneService(currentUser["user_username"])

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", currentUser["user_username"]),
		})
	}

	profileRsps := responses.GetProfileResponse(profile)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A username %s data are succesfully appeared!", currentUser["user_username"]),
		"data":    profileRsps,
	})
}

func (pc *profileController) UpdateProfileController(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)

	_, err := pc.profileService.FindOneService(currentUser["user_username"])

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", currentUser["user_username"]),
		})
	}

	var updateProfileInput inputs.UpdateProfileInput

	errUpdateProfileInput := c.BodyParser(&updateProfileInput)

	if errUpdateProfileInput != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   errUpdateProfileInput.Error(),
		})
	}

	updateProfile, _ := pc.profileService.UpdateService(currentUser["user_username"], updateProfileInput, currentUser)

	updateProfileRsps := responses.GetProfileResponse(updateProfile)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A username %s data are succesfully updated!", currentUser["user_username"]),
		"data":    updateProfileRsps,
	})
}

func (pc *profileController) UpdateBalanceTransactionController(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)

	_, err := pc.profileService.FindOneService(currentUser["user_username"])

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", currentUser["user_username"]),
		})
	}

	var updateProfileBalanceTransactionInput inputs.UpdateProfileBalanceTransactionInput

	errUpdateBalanceTransactionInput := c.BodyParser(&updateProfileBalanceTransactionInput)

	if errUpdateBalanceTransactionInput != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   errUpdateBalanceTransactionInput.Error(),
		})
	}

	updateBalanceTransaction, _ := pc.profileService.UpdateBalanceTransactionService(currentUser["user_username"], updateProfileBalanceTransactionInput, currentUser)

	updateBalanceTransactionRsps := responses.GetProfileBalanceTransactionResponse(updateBalanceTransaction)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A username %s balance transaction data are succesfully updated!", currentUser["user_username"]),
		"data":    updateBalanceTransactionRsps,
	})
}
