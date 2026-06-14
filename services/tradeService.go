package services

import (
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"slices"

	"github.com/google/uuid"
)

const tradesFile = "models/mock/trades.json"

func GetTrades() ([]*models.Trade, error) {
	trades, err := utils.ReadJSON[*models.Trade](tradesFile)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(trades, func(a, b *models.Trade) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})
	return trades, nil
}

func CreateTrade(newTrade models.Trade) error {
	if err := utils.UpsertJSON(tradesFile, &newTrade); err != nil {
		return err
	}
	return nil
}

func GetTradeForId(tradeId uuid.UUID) (*models.Trade, error) {
	trade, err := utils.GetJSON[*models.Trade](tradesFile, tradeId)
	if err != nil {
		return nil, err
	}
	return trade, nil
}

func UpdateTrade(trade models.Trade) error {
	if err := utils.UpsertJSON(tradesFile, &trade); err != nil {
		return err
	}
	return nil
}

func DeleteTrade(id uuid.UUID) error {
	if err := utils.DeleteJSON[*models.Trade](tradesFile, id); err != nil {
		return err
	}
	return nil
}
