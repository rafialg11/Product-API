package main

import (
	"product-api/database"
	"product-api/router"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":8080")
}
