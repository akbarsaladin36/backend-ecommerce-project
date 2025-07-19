package main

import (
	"backend-restapi-ecommerce/database"
	"backend-restapi-ecommerce/migration"
	"backend-restapi-ecommerce/router"
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	database.ConnectDB()
	migration.MigrateTables()
	router.ConnectRoutes()
}

func loadEnv() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load .env file")
	}
}
