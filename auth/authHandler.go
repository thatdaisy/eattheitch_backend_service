package auth

import (
	"eattheitch/backend/services"

	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func Login(context *gin.Context) {
	var req LoginRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := services.VerifyUserPassword(req.Email, req.Password); err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		log.Printf("invalid credentials: %s, %s", req.Email, req.Password)
		return
	}
	user, err := services.GetUserForEmail(req.Email)
	if err != nil {
		return
	}

	sessionID := createSession(user.Email)

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

	context.JSON(http.StatusOK, gin.H{"token": sessionID, "user": user})
}

func Current(context *gin.Context) {
	email := context.MustGet("email")
	emailString, _ := email.(string)
	user, _ := services.GetUserForEmail(emailString)
	context.JSON(http.StatusOK, gin.H{
		"message": "protected profile",
		"user":    user,
	})
}

func Logout(context *gin.Context) {
	sessionID, _ := context.Cookie("session_id")
	deleteSession(sessionID)
	context.SetCookie("session_id", "", -1, "/", "", false, true)
	context.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
