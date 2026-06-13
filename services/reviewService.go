package services

import (
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"encoding/json"
	"fmt"
	"log"
	"slices"

	"github.com/google/uuid"
)

const reviewsFile = "models/mock/reviews.json"

func GetReviews() ([]models.Review, error) {
	reviews, err := loadReviews()
	if err != nil {
		return nil, err
	}
	slices.SortFunc(reviews, func(a, b models.Review) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})
	return reviews, nil
}

func loadReviews() ([]models.Review, error) {
	data, err := utils.ReadJson(reviewsFile)
	if err != nil {
		log.Printf("could not read reviews from reviews.json - %s", err.Error())
		return []models.Review{}, nil
	}

	var raw []map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		fmt.Println("raw unmarshal error:", err)
	}
	for i, r := range raw {
		for _, key := range []string{"id", "brand_id"} {
			v, ok := r[key].(string)
			if !ok {
				fmt.Printf("record %d: field %q missing or not a string (value: %v)\n", i, key, r[key])
				continue
			}
			if _, err := uuid.Parse(v); err != nil {
				fmt.Printf("record %d: field %q = %q is invalid: %v\n", i, key, v, err)
			}
		}
	}

	var reviews []models.Review
	if err := json.Unmarshal(data, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}
