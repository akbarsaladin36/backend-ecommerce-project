package services

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/models"
	"backend-restapi-ecommerce/repositories"
	"time"
)

type UserService interface {
	FindAllService() ([]models.User, error)
	FindOneService(username string) (models.User, error)
	CreateService(userCreateInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error)
	UpdateService(username string, userUpdateInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error)
	DeleteService(username string) (models.User, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (us *userService) FindAllService() ([]models.User, error) {
	users, err := us.userRepository.FindAll()

	return users, err
}

func (us *userService) FindOneService(username string) (models.User, error) {
	user, err := us.userRepository.FindOne(username)

	return user, err
}

func (us *userService) CreateService(userCreateInput inputs.CreateUserInput, currentUser map[string]string) (models.User, error) {
	hashedPassword, _ := helper.HashPassword(userCreateInput.Password)
	userUUID := helper.GenerateUUID(userCreateInput.Username)

	user := models.User{
		UserUUID:            userUUID,
		UserUsername:        userCreateInput.Username,
		UserEmail:           userCreateInput.Email,
		UserPassword:        hashedPassword,
		UserFirstName:       userCreateInput.FirstName,
		UserLastName:        userCreateInput.LastName,
		UserAddress:         userCreateInput.Address,
		UserPhoneNumber:     userCreateInput.PhoneNumber,
		UserRole:            "user",
		UserStatusCd:        "active",
		UserCreatedDate:     time.Now(),
		UserCreatedUserUuid: currentUser["user_uuid"],
		UserCreatedUsername: currentUser["user_username"],
	}

	newUser, err := us.userRepository.Create(user)

	return newUser, err
}

func (us *userService) UpdateService(username string, userUpdateInput inputs.UpdateUserInput, currentUser map[string]string) (models.User, error) {
	checkUser, _ := us.userRepository.FindOne(username)

	checkUser.UserFirstName = userUpdateInput.FirstName
	checkUser.UserLastName = userUpdateInput.LastName
	checkUser.UserAddress = userUpdateInput.Address
	checkUser.UserPhoneNumber = userUpdateInput.PhoneNumber
	checkUser.UserUpdatedDate = time.Now()
	checkUser.UserUpdatedUserUuid = currentUser["user_uuid"]
	checkUser.UserUpdatedUsername = currentUser["user_username"]

	updateUser, err := us.userRepository.Update(checkUser)

	return updateUser, err
}

func (us *userService) DeleteService(username string) (models.User, error) {
	checkUser, _ := us.userRepository.FindOne(username)

	deleteUser, err := us.userRepository.Delete(checkUser)

	return deleteUser, err
}
