package handler

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewsResponse struct {
	Reviews []models.Review `json:"reviews"`
}

type NewReviewRequest struct {
	BrandID uuid.UUID `json:"brand_id" binding:"required"`
	Author  string    `json:"author" binding:"required"`

	RatingOverall    int `json:"rating_overall" binding:"required"`
	RatingSoftness   int `json:"rating_softness" binding:"required"`
	RatingQuality    int `json:"rating_quality" binding:"required"`
	RatingPriceValue int `json:"rating_price_value" binding:"required"`
	RatingEco        int `json:"rating_eco" binding:"required"`

	Text string `json:"text"`
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

func CreateReview(context *gin.Context) {
	var req NewReviewRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

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
