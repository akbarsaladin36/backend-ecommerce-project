package controllers

import (
	"backend-restapi-ecommerce/inputs"
	"backend-restapi-ecommerce/middleware"
	"backend-restapi-ecommerce/responses"
	"backend-restapi-ecommerce/services"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type invoiceController struct {
	invoiceService services.InvoiceService
}

func NewInvoiceController(invoiceService services.InvoiceService) *invoiceController {
	return &invoiceController{invoiceService}
}

func (ic *invoiceController) FindAllInvoices(c *fiber.Ctx) error {
	invoices, _ := ic.invoiceService.FindAllService()

	if len(invoices) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All invoices are empty!",
		})
	}

	var invoicesRsps []responses.InvoiceResponse

	for _, invoice := range invoices {
		invoiceRsps := responses.GetInvoiceResponse(invoice)

		invoicesRsps = append(invoicesRsps, invoiceRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All invoices are succesfully appeared!",
		"data":    invoicesRsps,
	})
}

func (ic *invoiceController) FindAllInvoiceByUserId(c *fiber.Ctx) error {
	currentUser := middleware.CurrentUser(c)

	invoices, _ := ic.invoiceService.FindAllByUserIdService(currentUser["user_uuid"])

	if len(invoices) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "All invoices for are empty!",
		})
	}

	var invoicesRsps []responses.InvoiceResponse

	for _, invoice := range invoices {
		invoiceRsps := responses.GetInvoiceResponse(invoice)

		invoicesRsps = append(invoicesRsps, invoiceRsps)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "All invoices for user are succesfully appeared!",
		"data":    invoicesRsps,
	})
}

func (ic *invoiceController) FindInvoiceController(c *fiber.Ctx) error {
	invoiceNo := c.Params("invoiceNo")

	invoice, err := ic.invoiceService.FindOneService(invoiceNo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A invoice number %s data are not found!", invoiceNo),
		})
	}

	invoiceRsps := responses.GetInvoiceResponse(invoice)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A invoice number %s data are succesfully appeared!", invoiceNo),
		"data":    invoiceRsps,
	})
}

func (ic *invoiceController) CreateInvoiceController(c *fiber.Ctx) error {
	var createInvoiceInput inputs.CreateInvoiceInput

	err := c.BodyParser(&createInvoiceInput)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	checkCart, errCheckCart := ic.invoiceService.FindCartService(createInvoiceInput.CartCode)

	if errCheckCart != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A cart code %s data are not found!", checkCart.CartCode),
		})
	}

	if checkCart.CartStatusCd == "paid" || checkCart.CartStatusCd == "paying" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A cart code %s data are paid and can't used again!", checkCart.CartCode),
		})
	}

	currentUser := middleware.CurrentUser(c)

	checkUser, errCheckUser := ic.invoiceService.FindUserService(currentUser["user_uuid"])

	if errCheckUser != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are not found!", checkUser.UserUsername),
		})
	}

	userBalanceTransactionAmount, _ := strconv.ParseInt(checkUser.UserBalanceTransactionAmount, 0, 64)

	if userBalanceTransactionAmount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A username %s data are have insufficient funds!", checkUser.UserUsername),
		})
	}

	invoicePrice, _ := strconv.ParseInt(createInvoiceInput.InvoicePrice, 0, 64)

	if userBalanceTransactionAmount < invoicePrice {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Your balance amount for transaction is less than your cart price. Please top up your balance first!",
		})
	}

	createInvoice, _ := ic.invoiceService.CreateService(createInvoiceInput, currentUser)

	createInvoiceRsps := responses.GetCreateInvoiceResponse(createInvoice)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "A new invoice are succesfully created!",
		"data":    createInvoiceRsps,
	})

}

func (ic *invoiceController) UpdateInvoiceController(c *fiber.Ctx) error {
	invoiceNo := c.Params("invoiceNo")

	_, err := ic.invoiceService.FindOneService(invoiceNo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A invoice number %s data are not found!", invoiceNo),
		})
	}

	var updateInvoiceInput inputs.UpdateInvoiceInput

	errUpdateInput := c.BodyParser(&updateInvoiceInput)

	if errUpdateInput != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "Invalid request body",
			"error":   errUpdateInput.Error(),
		})
	}

	currentUser := middleware.CurrentUser(c)

	updateInvoice, _ := ic.invoiceService.UpdateService(invoiceNo, updateInvoiceInput, currentUser)

	updateInvoiceRsps := responses.GetUpdateInvoiceResponse(updateInvoice)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A invoice number %s data are succesfully updated!", invoiceNo),
		"data":    updateInvoiceRsps,
	})

}

func (ic *invoiceController) DeleteInvoiceController(c *fiber.Ctx) error {
	invoiceNo := c.Params("invoiceNo")

	_, err := ic.invoiceService.FindOneService(invoiceNo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": fmt.Sprintf("A invoice number %s data are not found!", invoiceNo),
		})
	}

	ic.invoiceService.DeleteService(invoiceNo)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": fmt.Sprintf("A invoice number %s data are succesfully deleted!", invoiceNo),
	})
}
