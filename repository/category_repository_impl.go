package repository

import (
	"context"
	"database/sql"
	"errors"
	"kykurniawan/go-restful-api/helper"
	"kykurniawan/go-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
	//
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repositoru *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO categories (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(id)

	return category
}

func (repositoru *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE categories name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repositoru *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	SQL := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Id)
	helper.PanicIfError(err)
}

func (repositoru *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name FROM categories WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()
	category := domain.Category{}

	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		return category, nil
	} else {
		return category, errors.New("category is not found")
	}
}

func (repositoru *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, SQL)
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
