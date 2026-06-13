package handler

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReviewsResponse struct {
	Reviews []models.Review `json:"reviews"`
}

func GetReviews(context *gin.Context) {
	reviews, err := services.GetReviews()

	if err != nil {
		log.Printf("ERROR LoadingReviews %s", err)
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	var response ReviewsResponse
	response.Reviews = reviews
	context.JSON(http.StatusOK, response)
}
