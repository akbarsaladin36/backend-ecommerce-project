package router

import (
	"backend-restapi-ecommerce/controllers"
	"backend-restapi-ecommerce/database"
	"backend-restapi-ecommerce/middleware"
	"backend-restapi-ecommerce/repositories"
	"backend-restapi-ecommerce/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func ConnectRoutes() {
	app := fiber.New()
	app.Use(cors.New())

	authRoutes(app)
	adminRoutes(app)
	userRoutes(app)

	app.Listen(os.Getenv("APP_PORT"))
}

func authRoutes(app *fiber.App) {
	authRepositories := repositories.NewAuthRepository(database.DB)
	authServices := services.NewAuthService(authRepositories)
	authControllers := controllers.NewAuthController(authServices)

	v1Api := app.Group("/api/v1")

	v1Api.Post("/auth/register", authControllers.RegisterController)
	v1Api.Post("/auth/login", authControllers.LoginController)
}

func adminRoutes(app *fiber.App) {
	userRepositores := repositories.NewUserRepository(database.DB)
	userServices := services.NewUserService(userRepositores)
	userControllers := controllers.NewUserController(userServices)

	productRepositories := repositories.NewProductRepository(database.DB)
	productServices := services.NewProductService(productRepositories)
	productControllers := controllers.NewProductController(productServices)

	cartRepositories := repositories.NewCartRepository(database.DB)
	cartServices := services.NewCartService(cartRepositories)
	cartControllers := controllers.NewCartController(cartServices)

	invoiceRepositories := repositories.NewInvoiceRepository(database.DB)
	invoiceServices := services.NewInvoiceService(invoiceRepositories)
	invoiceControllers := controllers.NewInvoiceController(invoiceServices)

	profileRepositories := repositories.NewProfileRepository(database.DB)
	profileServices := services.NewProfileService(profileRepositories)
	profileControllers := controllers.NewProfileController(profileServices)

	v1Api := app.Group("/api/v1")
	v1Admin := v1Api.Group("/admin", middleware.AuthMiddleware, middleware.IsAdminAccess)

	// Users
	v1Admin.Get("/users", userControllers.FindUsersController)
	v1Admin.Get("/users/detail-user/:username", userControllers.FindUserController)
	v1Admin.Post("/users", userControllers.CreateUserController)
	v1Admin.Patch("/users/detail-user/:username", userControllers.UpdateUserController)
	v1Admin.Delete("/users/detail-user/:username", userControllers.DeleteUserController)

	// Products
	v1Admin.Get("/products", productControllers.FindProductsController)
	v1Admin.Get("/products/detail-product/:product_code", productControllers.FindProductController)
	v1Admin.Post("/products", productControllers.CreateProductController)
	v1Admin.Patch("/products/detail-product/:product_code", productControllers.UpdateProductController)
	v1Admin.Delete("/products/detail-product/:product_code", productControllers.DeleteProductController)

	// Carts
	v1Admin.Get("/carts", cartControllers.FindCartsController)
	v1Admin.Get("/carts/detail-cart/:cart_code", cartControllers.FindCartController)
	v1Admin.Post("/carts", cartControllers.CreateCartController)
	v1Admin.Patch("/carts/detail-cart/:cart_code", cartControllers.UpdateCartController)
	v1Admin.Delete("/carts/detail-cart/:cart_code", cartControllers.DeleteCartController)

	// Invoices
	v1Admin.Get("/invoices", invoiceControllers.FindAllInvoices)
	v1Admin.Get("/invoices/detail-invoice/:invoiceNo", invoiceControllers.FindInvoiceController)
	v1Admin.Post("/invoices", invoiceControllers.CreateInvoiceController)
	v1Admin.Patch("/invoices/detail-invoice/:invoiceNo", invoiceControllers.UpdateInvoiceController)
	v1Admin.Delete("/invoices/detail-invoice/:invoiceNo", invoiceControllers.DeleteInvoiceController)

	v1Admin.Get("/profile", profileControllers.FindProfileController)
	v1Admin.Patch("/profile", profileControllers.UpdateProfileController)
	v1Admin.Patch("/profile/update-balance-transaction", profileControllers.UpdateBalanceTransactionController)

}

func userRoutes(app *fiber.App) {
	v1Api := app.Group("/api/v1")
	v1User := v1Api.Group("/user", middleware.AuthMiddleware, middleware.IsUserAccess)

	productRepositories := repositories.NewProductRepository(database.DB)
	productServices := services.NewProductService(productRepositories)
	productControllers := controllers.NewProductController(productServices)

	cartRepositories := repositories.NewCartRepository(database.DB)
	cartServices := services.NewCartService(cartRepositories)
	cartControllers := controllers.NewCartController(cartServices)

	invoiceRepositories := repositories.NewInvoiceRepository(database.DB)
	invoiceServices := services.NewInvoiceService(invoiceRepositories)
	invoiceControllers := controllers.NewInvoiceController(invoiceServices)

	profileRepositories := repositories.NewProfileRepository(database.DB)
	profileServices := services.NewProfileService(profileRepositories)
	profileControllers := controllers.NewProfileController(profileServices)

	v1User.Get("/products", productControllers.FindProductsController)
	v1User.Get("/products/detail-product/:product_code", productControllers.FindProductController)

	v1User.Get("/carts", cartControllers.FindCartsByUserIdController)
	v1User.Get("/carts/detail-cart/:cart_code", cartControllers.FindCartController)
	v1User.Post("/carts", cartControllers.CreateCartController)
	v1User.Delete("/carts/detail-cart/:cart_code", cartControllers.DeleteCartController)

	v1User.Get("/invoices", invoiceControllers.FindAllInvoiceByUserId)
	v1User.Get("/invoices/detail-invoice/:invoiceNo", invoiceControllers.FindInvoiceController)
	v1User.Post("/invoices", invoiceControllers.CreateInvoiceController)

	v1User.Get("/profile", profileControllers.FindProfileController)
	v1User.Patch("/profile", profileControllers.UpdateProfileController)
	v1User.Patch("/profile/update-balance-transaction", profileControllers.UpdateBalanceTransactionController)
}
