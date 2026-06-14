package auth

import (
	"eattheitch/backend/handler"
	"eattheitch/backend/services"

	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
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

	session := sessions.Default(context)
	session.Set("email", req.Email)

	if err := session.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save session",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "logged in"})
}

func Current(context *gin.Context) {
	email := context.GetString("email")
	user, err := services.GetUserForEmail(email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to find user",
		})
		return
	}

	response := handler.SecureUserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Location:  user.Location,
		CreatedAt: user.CreatedAt,
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "protected profile",
		"user":    response,
	})
}

func Logout(context *gin.Context) {
	session := sessions.Default(context)

	session.Clear()

	if err := session.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear session"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "logged out"})
}
