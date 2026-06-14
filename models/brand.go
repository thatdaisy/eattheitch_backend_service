package models

import (
	"time"

	"github.com/google/uuid"
)

type Brand struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	OriginCountry  string    `json:"origin_country"`
	Description    string    `json:"description"`
	Certifications []string  `json:"certifications"`

	EcoScore      float64 `json:"eco_score"`
	AvgSoftness   float64 `json:"avg_softness"`
	AvgQuality    float64 `json:"avg_quality"`
	AvgPriceValue float64 `json:"avg_price_value"`
	AvgRating     float64 `json:"avg_rating"`
	ReviewCount   int     `json:"review_count"`

	CreatedAt time.Time `json:"created_at"`
}

func (b *Brand) GetID() uuid.UUID { return b.ID }
