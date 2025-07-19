package controllers

import (
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/middleware"
	"backend-restapi-ecommerce/responses"
	"backend-restapi-ecommerce/services"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *userController {
	return &userController{userService}
}

func (uc *userController) FindUsersController(c *fiber.Ctx) error {
	users, err := uc.userService.FindAllService()

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All messages are empty!",
		})
	}

	var userResponse []responses.UserResponse

	for _, user := range users {
		usersRsps := responses.GetUserResponse(user)

		userResponse = append(userResponse, usersRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All messages are succesfully appeared!",
		"data":    userResponse,
	})
}

func (uc *userController) FindUserController(c *fiber.Ctx) error {
	username := c.Params("username")

	user, err := uc.userService.FindOneService(username)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", username),
		})
	}

	userRsps := responses.GetUserResponse(user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A username %s data are succesfully appeared!", username),
		"data":    userRsps,
	})
}

func (uc *userController) CreateUserController(c *fiber.Ctx) error {
	var createUserInput inputs.CreateUserInput

	err := c.BodyParser(&createUserInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckUser := uc.userService.FindOneService(createUserInput.Username)

	if errCheckUser == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "A username has been registered! Please try again!",
		})
	}

	currentUser := middleware.CurrentUser(c)

	newCreateUser, _ := uc.userService.CreateService(createUserInput, currentUser)

	createUserRsps := responses.GetCreateUserResponse(newCreateUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A new user is succesfully created!",
		"data":    createUserRsps,
	})
}

func (uc *userController) UpdateUserController(c *fiber.Ctx) error {
	username := c.Params("username")

	var userUpdateInput inputs.UpdateUserInput

	err := c.BodyParser(&userUpdateInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	_, errCheckUser := uc.userService.FindOneService(username)

	if errCheckUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", username),
		})
	}

	currentUser := middleware.CurrentUser(c)

	newUpdateUser, _ := uc.userService.UpdateService(username, userUpdateInput, currentUser)

	updateUserRsps := responses.GetUpdateUserResponse(newUpdateUser)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A username %s data are succesfully updated!", username),
		"data":    updateUserRsps,
	})
}

func (uc *userController) DeleteUserController(c *fiber.Ctx) error {
	username := c.Params("username")

	_, errCheckUser := uc.userService.FindOneService(username)

	if errCheckUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", username),
		})
	}

	uc.userService.DeleteService(username)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A username %s data are succesfully deleted!", username),
	})
}
