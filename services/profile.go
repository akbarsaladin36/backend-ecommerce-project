package services

import (
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/models"
	"backend-restapi-ecommerce/repositories"
	"strconv"
	"time"
)

type ProfileService interface {
	FindOneService(user_username string) (models.User, error)
	UpdateService(user_username string, updateProfileInput inputs.UpdateProfileInput, currentUser map[string]string) (models.User, error)
	UpdateBalanceTransactionService(user_username string, updateProfileInput inputs.UpdateProfileBalanceTransactionInput, currentUser map[string]string) (models.User, error)
}

type profileService struct {
	profileRepository repositories.ProfileRepository
}

func NewProfileService(profileRepository repositories.ProfileRepository) *profileService {
	return &profileService{profileRepository}
}

func (ps *profileService) FindOneService(user_username string) (models.User, error) {
	user, err := ps.profileRepository.FindOne(user_username)

	return user, err
}

func (ps *profileService) UpdateService(user_username string, updateProfileInput inputs.UpdateProfileInput, currentUser map[string]string) (models.User, error) {
	checkProfile, _ := ps.profileRepository.FindOne(user_username)

	checkProfile.UserFirstName = updateProfileInput.FirstName
	checkProfile.UserLastName = updateProfileInput.LastName
	checkProfile.UserAddress = updateProfileInput.Address
	checkProfile.UserPhoneNumber = updateProfileInput.PhoneNumber
	checkProfile.UserUpdatedDate = time.Now()
	checkProfile.UserUpdatedUserUuid = currentUser["user_uuid"]
	checkProfile.UserUpdatedUsername = currentUser["user_username"]

	updateProfile, err := ps.profileRepository.Update(checkProfile)

	return updateProfile, err
}

func (ps *profileService) UpdateBalanceTransactionService(user_username string, updateProfileBalanceTransaction inputs.UpdateProfileBalanceTransactionInput, currentUser map[string]string) (models.User, error) {
	checkProfile, _ := ps.profileRepository.FindOne(user_username)

	profileBalanceTransaction, _ := strconv.ParseInt(checkProfile.UserBalanceTransactionAmount, 0, 64)
	inputBalanceTransaction, _ := strconv.ParseInt(updateProfileBalanceTransaction.BalanceTransaction, 0, 64)
	calculateBalanceTransaction := profileBalanceTransaction + inputBalanceTransaction
	formattedBalanceTransactionResult := strconv.FormatInt(calculateBalanceTransaction, 10)

	checkProfile.UserBalanceTransactionAmount = formattedBalanceTransactionResult
	checkProfile.UserUpdatedDate = time.Now()
	checkProfile.UserUpdatedUserUuid = currentUser["user_uuid"]
	checkProfile.UserUpdatedUsername = currentUser["user_username"]

	updateProfile, err := ps.profileRepository.Update(checkProfile)

	return updateProfile, err
}
