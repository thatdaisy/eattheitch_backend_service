package handler

import (
	"eattheitch/backend/models"
	"eattheitch/backend/services"
	"eattheitch/backend/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BrandsResponse struct {
	Brands     []models.Brand `json:"brands"`
	Page       int            `json:"page"`
	Total      int            `json:"total"`
	TotalPages int            `json:"total_pages"`
}

func GetBrands(context *gin.Context) {
	pageQuery := context.DefaultQuery("page", "1")
	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		log.Printf("could not parse param page [%s]", pageQuery)
		page = 1
	}
	perPageQuery := context.DefaultQuery("per_page", "10")
	perPage, err := strconv.Atoi(perPageQuery)
	if err != nil {
		log.Printf("could not parse param per_page [%s]", pageQuery)
		page = 10
	}
	sortQuery := services.SortOrder(context.DefaultQuery("sort", ""))
	if !sortQuery.IsValid() {
		log.Printf("invalid sort_by value %s ", sortQuery)
		sortQuery = services.SortByNameAsc
	}

	brands, err := services.GetBrandsSorted(sortQuery)
	if err != nil {
		log.Printf("ERROR LoadingBrands %s", err)
		context.JSON(http.StatusInternalServerError, gin.H{

			"error": err.Error(),
		})
		return
	}

	items, total, totalPages := utils.Paginate(brands, page, perPage)
	var response BrandsResponse
	response.Brands = items
	response.Page = page
	response.Total = total
	response.TotalPages = totalPages

	context.JSON(http.StatusOK, response)
}
