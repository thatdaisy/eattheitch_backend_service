package main

import (
	"eattheitch/backend/api"
)

func main() {
	router := api.SetupRoutes()
	router.Run(":8080")
}
