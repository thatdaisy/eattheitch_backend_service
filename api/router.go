package api

import (
	"time"

	"eattheitch/backend/auth"
	"eattheitch/backend/handler"
	"eattheitch/backend/middleware"
	"eattheitch/backend/thirdparty"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	middleware.SessionSetup(router)

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
	router.GET("/countries", thirdparty.GetCountries)

	protected := router.Group("/")
	protected.Use(middleware.AuthRequired())
	{
		protected.GET("/auth/current", auth.Current)
		protected.POST("/auth/logout", auth.Logout)
		protected.PUT("/users/:userId", handler.UpdateUser)

		protected.GET("/brands", handler.GetBrands)
		protected.GET("/brands/:brandId", handler.GetBrandForId)

		protected.GET("/reviews", handler.GetReviews)
		protected.POST("/brands/:brandId/reviews", handler.CreateReview)
		protected.PUT("/reviews/:reviewId", handler.UpdateReview)
		protected.DELETE("/reviews/:reviewId", handler.DeleteReview)

		protected.GET("/trades", handler.GetTrades)
		protected.GET("/trades/:tradeId", handler.GetTradeForId)
		protected.POST("/trades", handler.CreateTrade)
		protected.PUT("/trades/:tradeId", handler.UpdateTrade)
		protected.DELETE("/trades/:tradeId", handler.DeleteTrade)
	}

	return router
}
