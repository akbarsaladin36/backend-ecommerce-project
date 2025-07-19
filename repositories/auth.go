package repositories

import (
	"backend-restapi-ecommerce/models"

	"gorm.io/gorm"
)

type AuthRepository interface {
	FindOne(username string) (models.User, error)
	Create(user models.User) (models.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{db}
}

func (ar *authRepository) FindOne(username string) (models.User, error) {
	var user models.User

	err := ar.db.Where("user_username = ?", username).First(&user).Error

	return user, err
}

func (ar *authRepository) Create(user models.User) (models.User, error) {
	err := ar.db.Create(&user).Error

	return user, err
}
