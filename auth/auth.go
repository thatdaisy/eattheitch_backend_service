package auth

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
	// Login
	router.POST("/auth/login", func(context *gin.Context) {
		var req models.User
		if err := context.ShouldBindJSON(&req); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		// validate user
		u, err := services.GetUserFormEmail(req.Email)
		log.Printf("found user: %s, %s, %s, %s", u.Email, u.Password, u.Username, u.Location)
		log.Printf("password match: %s == %s", u.Password, req.Password)
		if err != nil || u.Password != req.Password {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
			log.Printf("invalid credentials: %s, %s", req.Email, req.Password)
			return
		}

		sessionID := createSession(u.Username)

		// set cookie
		context.SetCookie(
			"session_id",
			sessionID,
			1800, // 30 minutes
			"/",
			"",
			false,
			true,
		)

		context.JSON(http.StatusOK, gin.H{"token": sessionID, "user": u})
	})

	router.GET("/auth/me", isAuhenticated(), func(context *gin.Context) {
		username := context.MustGet("username")
		context.JSON(http.StatusOK, gin.H{
			"message": "protected profile",
			"user":    username,
		})
	})

	router.POST("/auth/logout", isAuhenticated(), func(context *gin.Context) {
		sessionID, _ := context.Cookie("session_id")

		deleteSession(sessionID)

		context.SetCookie("session_id", "", -1, "/", "", false, true)

		context.JSON(http.StatusOK, gin.H{"message": "logged out"})
	})
}
