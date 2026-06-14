package models

import (
	"time"

	"github.com/google/uuid"
)

type Trade struct {
	ID        uuid.UUID `json:"id"`
	Author    string    `json:"author"`
	BrandName string    `json:"brand_name"`
	Location  string    `json:"location"`
	Title     string    `json:"title"`
	Text      string    `json:"text,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Trade) GetID() uuid.UUID { return t.ID }
