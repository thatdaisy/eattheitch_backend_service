package models

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ID      uuid.UUID `json:"id"`
	BrandID uuid.UUID `json:"brand_id"`
	Author  string    `json:"author"`

	RatingOverall    int `json:"rating_overall"`
	RatingSoftness   int `json:"rating_softness"`
	RatingQuality    int `json:"rating_quality"`
	RatingPriceValue int `json:"rating_price_value"`
	RatingEco        int `json:"rating_eco"`

	Text string `json:"text,omitempty"`

	CreatedAt time.Time `json:"created_at"`
}

func (r *Review) GetID() uuid.UUID { return r.ID }
