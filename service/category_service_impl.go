package service

import (
	"context"
	"database/sql"
	"golang-resful-api/exception"
	"golang-resful-api/helper"
	"golang-resful-api/model/domain"
	"golang-resful-api/model/web"
	"golang-resful-api/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(repository repository.CategoryRepository, Db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: repository,
		DB:                 Db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, req web.CategoryCreateRequest) (web.CategoryResponse, error) {

	// Validation
	errValidation := service.Validate.Struct(req)
	helper.PanicIfError(errValidation)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// Check the transaction
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: req.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)

	return helper.ConvertToCategoryResponse(category), nil
}

func (service *CategoryServiceImpl) Update(ctx context.Context, req web.CategoryUpdateRequest) (web.CategoryResponse, error) {
	// Validation error checking
	errValidation := service.Validate.Struct(req)
	helper.PanicIfError(errValidation)

	// Create Transaction First
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	// Check the transaction
	defer helper.CommitOrRollback(tx)

	// Check data category if is it exist or not
	response, err := service.CategoryRepository.FindById(ctx, tx, req.Id)

	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	category := domain.Category{
		Id:   response.Id,
		Name: req.Name,
	}

	// Update Category
	updatedCategory := service.CategoryRepository.Update(ctx, tx, category)

	return helper.ConvertToCategoryResponse(updatedCategory), nil
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	// Check data category is it exist or not
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.CategoryRepository.Delete(ctx, tx, category.Id)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ConvertToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.CategoryRepository.FindAll(ctx, tx)

	return helper.ConverToSliceCategoryResponse(categories)
}
