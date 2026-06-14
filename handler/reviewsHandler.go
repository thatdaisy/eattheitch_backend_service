package handler

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReviewsResponse struct {
	Reviews []models.Review `json:"reviews"`
}

type NewReviewRequest struct {
	Author string `json:"author" binding:"required"`

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
	brandId, err := uuid.Parse(context.Param("brandId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req NewReviewRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newReview := models.Review{
		ID:               uuid.New(),
		BrandID:          brandId,
		Author:           req.Author,
		RatingOverall:    req.RatingOverall,
		RatingSoftness:   req.RatingSoftness,
		RatingQuality:    req.RatingQuality,
		RatingPriceValue: req.RatingPriceValue,
		RatingEco:        req.RatingEco,
		Text:             req.Text,
		CreatedAt:        time.Now(),
	}

	services.CreateReview(newReview)

	context.JSON(http.StatusOK, newReview)
}
