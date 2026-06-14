package handler

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"eattheitch/backend/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NewReviewRequest struct {
	Author           string `json:"author" binding:"required"`
	RatingOverall    int    `json:"rating_overall" binding:"required"`
	RatingSoftness   int    `json:"rating_softness" binding:"required"`
	RatingQuality    int    `json:"rating_quality" binding:"required"`
	RatingPriceValue int    `json:"rating_price_value" binding:"required"`
	RatingEco        int    `json:"rating_eco" binding:"required"`
	Text             string `json:"text" binding:"max=500"`
}

type UpdateReviewRequest struct {
	BrandID          *uuid.UUID `json:"brand_id"`
	RatingOverall    *int       `json:"rating_overall"`
	RatingSoftness   *int       `json:"rating_softness"`
	RatingQuality    *int       `json:"rating_quality"`
	RatingPriceValue *int       `json:"rating_price_value"`
	RatingEco        *int       `json:"rating_eco"`
	Text             *string    `json:"text" binding:"max=500"`
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
	context.JSON(http.StatusOK, reviews)
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

func UpdateReview(context *gin.Context) {
	reviewId, err := uuid.Parse(context.Param("reviewId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req UpdateReviewRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateReview, err := services.GetReviewForId(reviewId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applyReviewUpdate(updateReview, req)
	utils.UpsertJSON("", updateReview)

	context.JSON(http.StatusOK, updateReview)
}

func applyReviewUpdate(review *models.Review, req UpdateReviewRequest) {
	utils.SetIfNotNil(&review.BrandID, req.BrandID)
	utils.SetIfNotNil(&review.RatingOverall, req.RatingOverall)
	utils.SetIfNotNil(&review.RatingSoftness, req.RatingSoftness)
	utils.SetIfNotNil(&review.RatingQuality, req.RatingQuality)
	utils.SetIfNotNil(&review.RatingPriceValue, req.RatingPriceValue)
	utils.SetIfNotNil(&review.RatingEco, req.RatingEco)
	utils.SetIfNotNil(&review.Text, req.Text)
}
