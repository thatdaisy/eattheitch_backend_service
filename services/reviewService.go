package services

import (
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"encoding/json"
	"log"
	"slices"
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

func CreateReview(newReview models.Review) error {
	reviews, err := loadReviews()
	if err != nil {
		return err
	}
	reviews = append(reviews, newReview)
	if err := utils.WriteJson(reviewsFile, reviews); err != nil {
		return err
	}
	return nil
}

func loadReviews() ([]models.Review, error) {
	data, err := utils.ReadJson(reviewsFile)
	if err != nil {
		log.Printf("could not read reviews from reviews.json - %s", err.Error())
		return []models.Review{}, nil
	}

	var reviews []models.Review
	if err := json.Unmarshal(data, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}
