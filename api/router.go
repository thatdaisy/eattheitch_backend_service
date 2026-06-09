package api

import (
	"eattheitch/backend/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	auth.SetupAuthRoutes(router)

	router.Run()
	return router
}
