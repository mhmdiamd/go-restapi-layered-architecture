package helper

import (
	"golang-resful-api/model/domain"
	"golang-resful-api/model/web"
)

func ConvertToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ConverToSliceCategoryResponse(categories []domain.Category) []web.CategoryResponse {
	var dataCategories []web.CategoryResponse
	for _, category := range categories {
		dataCategories = append(dataCategories, ConvertToCategoryResponse(category))
	}

	return dataCategories
}
