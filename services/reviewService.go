package services

import (
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"errors"
	"slices"

	"github.com/google/uuid"
)

const reviewsFile = "models/mock/reviews.json"

func GetReviews() ([]*models.Review, error) {
	reviews, err := utils.ReadJSON[*models.Review](reviewsFile)
	if err != nil {
		return nil, err
	}
	slices.SortFunc(reviews, func(a, b *models.Review) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})
	return reviews, nil
}

func CreateReview(newReview models.Review) error {
	if err := utils.UpsertJSON(reviewsFile, &newReview); err != nil {
		return err
	}
	return nil
}

func GetReviewForId(reviewId uuid.UUID) (*models.Review, error) {
	reviews, err := utils.ReadJSON[*models.Review](reviewsFile)
	if err != nil {
		return nil, err
	}
	for _, review := range reviews {
		if review.ID == reviewId {
			return review, nil
		}
	}
	return nil, errors.New("review not found " + reviewId.String())
}

func UpdateReview(review models.Review) error {
	if err := utils.UpsertJSON(reviewsFile, &review); err != nil {
		return err
	}
	return nil
}
