package handler

import (
	"eattheitch/backend/services"
	"eattheitch/backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SecureUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Location  string    `json:"location,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateUserRequest struct {
	Location string `json:"location"`
}

func UpdateUser(context *gin.Context) {
	userId, err := uuid.Parse(context.Param("userId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req UpdateUserRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateUser, err := services.GetUserForId(userId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	utils.SetIfNotNil(&updateUser.Location, &req.Location)

	if err := services.UpdateUser(*updateUser); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := SecureUserResponse{
		ID:        updateUser.ID,
		Email:     updateUser.Email,
		Username:  updateUser.Username,
		Location:  updateUser.Location,
		CreatedAt: updateUser.CreatedAt,
	}
	context.JSON(http.StatusOK, response)
}
