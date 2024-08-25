package main

import (
	"natural-products-api/database"
	"natural-products-api/routes"
)

func main() {
	database.ConnectDatabase()

	router := routes.SetupRouter()

	router.Run(":8080")
}
