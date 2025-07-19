package helper

import (
	"backend-restapi-ecommerce/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	UserUuid     string `json:"user_uuid"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.User) (tokenString string, err error) {
	claims := &JWTClaim{
		UserUuid:     user.UserUUID,
		UserUsername: user.UserUsername,
		UserEmail:    user.UserEmail,
		UserRole:     user.UserRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}
