package handler

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"eattheitch/backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type NewTradeRequest struct {
	Author    string `json:"author"`
	BrandName string `json:"brand_name"`
	Location  string `json:"location"`
	Title     string `json:"title"`
	Text      string `json:"text,omitempty"`
}

type UpdateTradeRequest struct {
	BrandName string `json:"brand_name"`
	Location  string `json:"location"`
	Title     string `json:"title"`
	Text      string `json:"text,omitempty"`
}

func GetTrades(context *gin.Context) {
	trades, err := services.GetTrades()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, trades)
}

func GetTradeForId(context *gin.Context) {
	tradeId, err := uuid.Parse(context.Param("tradeId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	trade, err := services.GetTradeForId(tradeId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	context.JSON(http.StatusOK, trade)
}

func CreateTrade(context *gin.Context) {
	var req NewTradeRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	newTrade := models.Trade{
		ID:        uuid.New(),
		Author:    req.Author,
		BrandName: req.BrandName,
		Location:  req.Location,
		Title:     req.Title,
		Text:      req.Text,
		CreatedAt: time.Now(),
	}

	services.CreateTrade(newTrade)
	context.JSON(http.StatusOK, newTrade)
}

func UpdateTrade(context *gin.Context) {
	tradeId, err := uuid.Parse(context.Param("tradeId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var req UpdateTradeRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateTrade, err := services.GetTradeForId(tradeId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applyTradeUpdate(updateTrade, req)

	if err := services.UpdateTrade(*updateTrade); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, updateTrade)
}

func applyTradeUpdate(trade *models.Trade, req UpdateTradeRequest) {
	utils.SetIfNotNil(&trade.BrandName, &req.BrandName)
	utils.SetIfNotNil(&trade.Location, &req.Location)
	utils.SetIfNotNil(&trade.Title, &req.Title)
	utils.SetIfNotNil(&trade.Text, &req.Text)
}

func DeleteTrade(context *gin.Context) {
	tradeId, err := uuid.Parse(context.Param("tradeId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	services.DeleteTrade(tradeId)
}
