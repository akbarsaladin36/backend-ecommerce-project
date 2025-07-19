package migration

import (
	"backend-restapi-ecommerce/database"
	"backend-restapi-ecommerce/models"
	"fmt"
)

func MigrateTables() {
	migrateUser()
	migrateProduct()
	migrateCart()
	migrateInvoice()
}

func migrateUser() {

	// database.DB.Migrator().DropTable(models.User{})

	if !database.DB.Migrator().HasTable(models.User{}) {
		err := database.DB.AutoMigrate(models.User{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "User")
	}
}

func migrateProduct() {

	// database.DB.Migrator().DropTable(models.Product{})

	if !database.DB.Migrator().HasTable(models.Product{}) {
		err := database.DB.AutoMigrate(models.Product{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Product")
	}
}

func migrateCart() {

	// database.DB.Migrator().DropTable(models.Cart{})

	if !database.DB.Migrator().HasTable(models.Cart{}) {
		err := database.DB.AutoMigrate(models.Cart{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Cart")
	}
}

func migrateInvoice() {

	// database.DB.Migrator().DropTable(models.Invoice{})

	if !database.DB.Migrator().HasTable(models.Invoice{}) {
		err := database.DB.AutoMigrate(models.Invoice{})
		if err != nil {
			fmt.Println("Error migrating table", err)
			return
		}
		fmt.Printf("Table %s is succesfully migrated!\n", "Invoice")
	}
}
