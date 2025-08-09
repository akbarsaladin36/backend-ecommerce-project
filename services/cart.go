package services

import (
	"backend-restapi-ecommerce/helper"
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/models"
	"backend-restapi-ecommerce/repositories"
	"strconv"
	"time"
)

type CartService interface {
	FindAllService() ([]models.CartWithProduct, error)
	FindAllByUserIdService(user_uuid string) ([]models.CartWithProduct, error)
	FindOneService(cart_code string) (models.Cart, error)
	CreateService(createCartInput inputs.CreateCartInput, currentUser map[string]string) (models.Cart, error)
	UpdateService(cart_code string, updateCartInput inputs.UpdateCartInput, currentUser map[string]string) (models.Cart, error)
	DeleteService(cart_code string) (models.Cart, error)
}

type cartService struct {
	cartRepository repositories.CartRepository
}

func NewCartService(cartRepository repositories.CartRepository) *cartService {
	return &cartService{cartRepository}
}

func (cs *cartService) FindAllService() ([]models.CartWithProduct, error) {
	carts, err := cs.cartRepository.FindAll()

	return carts, err
}

func (cs *cartService) FindAllByUserIdService(user_uuid string) ([]models.CartWithProduct, error) {
	carts, err := cs.cartRepository.FindAllByUserId(user_uuid)

	return carts, err
}

func (cs *cartService) FindOneService(cart_code string) (models.Cart, error) {
	cart, err := cs.cartRepository.FindOne(cart_code)

	return cart, err
}

func (cs *cartService) CreateService(createCartInput inputs.CreateCartInput, currentUser map[string]string) (models.Cart, error) {
	cartCode := helper.GenerateUUID(createCartInput.ProductCode)

	checkProduct, _ := cs.cartRepository.FindProduct(createCartInput.ProductCode)

	productPrice, _ := strconv.ParseInt(checkProduct.ProductPrice, 0, 64)
	cartQuantity, _ := strconv.ParseInt(createCartInput.CartQuantity, 0, 64)
	totalPrice := productPrice * cartQuantity
	totalPriceToString := strconv.FormatInt(totalPrice, 10)

	cart := models.Cart{
		UserUuid:            currentUser["user_uuid"],
		ProductCode:         createCartInput.ProductCode,
		CartCode:            cartCode,
		CartDescription:     createCartInput.CartDescription,
		CartPrice:           totalPriceToString,
		CartQuantity:        createCartInput.CartQuantity,
		CartStatusCd:        "pending",
		CartCreatedDate:     time.Now(),
		CartCreatedUserUuid: currentUser["user_uuid"],
		CartCreatedUsername: currentUser["user_username"],
	}

	newCart, err := cs.cartRepository.Create(cart)

	return newCart, err
}

func (cs *cartService) UpdateService(cart_code string, updateCartInput inputs.UpdateCartInput, currentUser map[string]string) (models.Cart, error) {
	checkCart, _ := cs.cartRepository.FindOne(cart_code)
	checkProduct, _ := cs.cartRepository.FindProduct(updateCartInput.ProductCode)

	productPrice, _ := strconv.ParseInt(checkProduct.ProductPrice, 0, 64)
	cartQuantity, _ := strconv.ParseInt(updateCartInput.CartQuantity, 0, 64)
	totalPrice := productPrice * cartQuantity
	totalPriceToString := strconv.FormatInt(totalPrice, 10)

	checkCart.ProductCode = updateCartInput.ProductCode
	checkCart.CartDescription = updateCartInput.CartDescription
	checkCart.CartPrice = totalPriceToString
	checkCart.CartQuantity = updateCartInput.CartQuantity
	checkCart.CartUpdatedDate = time.Now()
	checkCart.CartUpdatedUserUuid = currentUser["user_uuid"]
	checkCart.CartUpdatedUsername = currentUser["user_username"]

	updateCart, err := cs.cartRepository.Update(checkCart)

	return updateCart, err
}

func (cs *cartService) DeleteService(cart_code string) (models.Cart, error) {
	checkCart, _ := cs.cartRepository.FindOne(cart_code)

	deleteCart, err := cs.cartRepository.Delete(checkCart)

	return deleteCart, err
}
