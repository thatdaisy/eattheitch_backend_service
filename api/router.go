package api

import (
	"time"

	"eattheitch/backend/auth"
	"eattheitch/backend/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.SetTrustedProxies([]string{})

	router.POST("/auth/register", auth.Register)
	router.POST("/auth/login", auth.Login)
	router.GET("/auth/current", auth.IsAuhenticated(), auth.Current)
	router.POST("/auth/logout", auth.IsAuhenticated(), auth.Logout)

	router.GET("/brands", auth.IsAuhenticated(), handler.GetBrands)

	router.GET("/reviews", auth.IsAuhenticated(), handler.GetReviews)
	router.POST("/brands/:brandId/reviews", auth.IsAuhenticated(), handler.CreateReview)

	return router
}
