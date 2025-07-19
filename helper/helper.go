package helper

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID(data string) (newUUIDString string) {
	namespace := uuid.NameSpaceDNS
	customUUID := uuid.NewSHA1(namespace, []byte(data))
	uuidString := customUUID.String()
	uuidString = strings.ReplaceAll(uuidString, "-", "")

	return uuidString
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPassword(user_password string, input_password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user_password), []byte(input_password))

	if err != nil {
		fmt.Println("Your password is not matched! Please try again!")
	}

	return err
}

func GenerateSlug(input string) string {
	slug := slug.Make(input)
	return slug
}

func GenerateInvoice() string {
	now := time.Now()
	randomInt := strconv.Itoa(rand.Intn(100000))

	formatInvoice := fmt.Sprintf("%s-%04d%02d%02d-%s", "INV", now.Year(), now.Month(), now.Day(), randomInt)

	return formatInvoice
}
