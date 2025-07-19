package middleware

import (
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_KEY"))

type AuthMiddlewareClaim struct {
	UserUuid     string `json:"user_uuid"`
	UserUsername string `json:"user_username"`
	UserEmail    string `json:"user_email"`
	UserRole     string `json:"user_role"`
	jwt.RegisteredClaims
}

func AuthMiddleware(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")

	if tokenString == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Token is not found / empty!",
			"status":  fiber.StatusNotFound,
		})

	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Invalid token format!",
		})

	}

	claims := &AuthMiddlewareClaim{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrTokenExpired {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":  fiber.StatusUnauthorized,
				"message": "Token is expired!",
			})
		}

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Token tidak berhasil digenerate!",
		})
	}

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Token tidak sesuai!",
		})
	}

	if time.Now().After(claims.ExpiresAt.Time) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Token is expired!",
		})
	}

	currentUser := map[string]string{
		"user_uuid":     claims.UserUuid,
		"user_username": claims.UserUsername,
		"user_email":    claims.UserEmail,
		"user_role":     claims.UserRole,
	}

	c.Locals("currentUser", currentUser)

	return c.Next()
}

func IsAdminAccess(c *fiber.Ctx) error {
	getCurrentUser := c.Locals("currentUser").(map[string]string)

	if getCurrentUser["user_role"] != "admin" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "This url can be accessed by admin",
		})
	}

	return c.Next()
}

func IsUserAccess(c *fiber.Ctx) error {
	getCurrentUser := c.Locals("currentUser").(map[string]string)

	if getCurrentUser["user_role"] != "user" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "This url can be accessed by user",
		})
	}

	return c.Next()
}

func CurrentUser(c *fiber.Ctx) map[string]string {
	getCurrentUser := c.Locals("currentUser").(map[string]string)

	return getCurrentUser
}
