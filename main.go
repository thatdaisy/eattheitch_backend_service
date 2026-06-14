package main

import (
	"eattheitch/backend/api"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := api.SetupRoutes()
	router.Run(":8080")
}
