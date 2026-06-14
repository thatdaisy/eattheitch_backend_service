package auth

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=30"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func Register(context *gin.Context) {
	var req RegisterRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	checkEmail, _ := services.GetUserForEmail(req.Email)
	checkUsername, _ := services.GetUserForUsername(req.Email)
	if checkEmail != nil || checkUsername != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "user already exist",
		})
		return
	}

	passwordHash, err := services.HashUserPassword(req.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to process password",
		})
		return
	}

	newUser := models.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		CreatedAt:    time.Now(),
	}

	if err := services.SaveUser(&newUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to save user",
		})
		return
	}

	responseUser := newUser
	responseUser.PasswordHash = ""

	context.JSON(http.StatusCreated, AuthResponse{
		User: responseUser,
	})
}
