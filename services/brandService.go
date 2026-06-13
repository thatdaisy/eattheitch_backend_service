package services

import (
	"cmp"
	"eattheitch/backend/models"
	"eattheitch/backend/utils"
	"encoding/json"
	"log"
	"slices"
)

const brandsFile = "models/mock/brands.json"

type SortOrder string

const (
	SortByNameAsc     SortOrder = "name_asc"
	SortByNameDesc    SortOrder = "name_desc"
	SortAvgRatingAsc  SortOrder = "avg_rating_asc"
	SortAvgRatingDesc SortOrder = "avg_rating_desc"
	SortEcoScoreAsc   SortOrder = "eco_score_asc"
	SortEcoScoreDesc  SortOrder = "eco_score_desc"
)

func (s SortOrder) IsValid() bool {
	switch s {
	case SortByNameAsc, SortByNameDesc, SortAvgRatingAsc, SortAvgRatingDesc, SortEcoScoreAsc, SortEcoScoreDesc:
		return true
	}
	return false
}

func GetBrandsSorted(sort SortOrder) ([]models.Brand, error) {
	brands, err := loadBrands()
	if err != nil {
		return nil, err
	}

	switch sort {
	case SortByNameAsc:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			return cmp.Compare(a.Name, b.Name)
		})
	case SortByNameDesc:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			return cmp.Compare(b.Name, a.Name)
		})
	case SortAvgRatingAsc:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			if a.AvgRating != b.AvgRating {
				return cmp.Compare(a.AvgRating, b.AvgRating)
			}
			return cmp.Compare(a.Name, b.Name)
		})
	case SortAvgRatingDesc:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			if a.AvgRating != b.AvgRating {
				return cmp.Compare(b.AvgRating, a.AvgRating)
			}
			return cmp.Compare(a.Name, b.Name)
		})
	case SortEcoScoreAsc:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			if a.EcoScore != b.EcoScore {
				return cmp.Compare(a.EcoScore, b.EcoScore)
			}
			return cmp.Compare(a.Name, b.Name)
		})
	case SortEcoScoreDesc:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			if a.EcoScore != b.EcoScore {
				return cmp.Compare(b.EcoScore, a.EcoScore)
			}
			return cmp.Compare(a.Name, b.Name)
		})
	default:
		slices.SortFunc(brands, func(a, b models.Brand) int {
			if a.AvgRating != b.AvgRating {
				return cmp.Compare(a.AvgRating, b.AvgRating)
			}
			return cmp.Compare(a.Name, b.Name)
		})
	}
	return brands, nil
}

func loadBrands() ([]models.Brand, error) {
	data, err := utils.ReadJson(brandsFile)
	if err != nil {
		log.Printf("could not read brands from brands.json - %s", err.Error())
		return []models.Brand{}, nil
	}

	var brands []models.Brand
	if err := json.Unmarshal(data, &brands); err != nil {
		return nil, err
	}
	return brands, nil
}
