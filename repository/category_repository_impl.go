package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"golang-resful-api/helper"
	"golang-resful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	query := `INSERT INTO category VALUES(DEFAULT, $1) RETURNING id`
	rows := tx.QueryRow(query, category.Name)

	var data domain.Category
	var dataId int
	err := rows.Scan(&dataId)
	helper.PanicIfError(err)

	data = domain.Category{
		Id:   dataId,
		Name: category.Name,
	}

	fmt.Println(data)

	return data
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {

	fmt.Println(category)
	query := "UPDATE category SET name = $1 WHERE id = $2"
	rows, err := tx.QueryContext(ctx, query, category.Name, category.Id)

	// Chech if Error will throw panic (Error handling)
	helper.PanicIfError(err)
	defer rows.Close()

	fmt.Println(category)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	query := "DELETE FROM category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)

	// Chech if Error will throw panic (Error handling)
	helper.PanicIfError(err)
	defer rows.Close()

}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	query := "SELECT * from category WHERE id = $1"
	rows, err := tx.QueryContext(ctx, query, categoryId)

	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT * from category"
	rows, err := tx.QueryContext(ctx, query)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category

	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)

		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
