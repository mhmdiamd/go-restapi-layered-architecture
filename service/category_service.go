package service

import (
	"context"
	"golang-resful-api/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryCreateRequest) (web.CategoryResponse, error)
	Update(ctx context.Context, request web.CategoryUpdateRequest) (web.CategoryResponse, error)
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}
