package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"go-restaurant-api/api/config/database"
	"go-restaurant-api/api/config/database/statement"
	"go-restaurant-api/api/enum"
	"go-restaurant-api/api/model/entity"
	"go-restaurant-api/api/model/exception"
	"log"
)

var CategoryRepo categoryRepository

type categoryRepository struct {
	stmt *statement.CategoryStatement
}

func init() {
	CategoryRepo = categoryRepository{
		stmt: &statement.CategoryStmt,
	}
}

func (repository *categoryRepository) ListAllCategories() []entity.Category {
	rows, err := repository.stmt.FindAll.Query()
	defer closeRows(rows, enum.CategoryNotFound)
	if err != nil {
		log.Panicln(err)
	}
	return loopThroughCategoryRows(rows)
}

func (repository *categoryRepository) GetCategory(categoryId int64) *entity.Category {
	row := repository.stmt.FindById.QueryRow(categoryId)
	return rowToCategory(row)
}

func (repository *categoryRepository) CreateCategory(category *entity.Category) {
	log.Printf("Creating category. Name: %s.", category.CategoryName)
	err := repository.stmt.Create.QueryRow(category.CategoryName).Scan(&category.Id)
	checkUniqueException(enum.CategoryNotCreated, err)
	log.Printf("New category. Id: %d, Name: %s.", category.Id, category.CategoryName)
}

func (repository *categoryRepository) UpdateCategory(category *entity.Category) {
	log.Printf("Updating category. Id: %d.", category.Id)
	_, err := repository.stmt.Update.Exec(category.CategoryName, category.Id)
	checkUniqueException(enum.CategoryNotUpdate, err)
	log.Printf("Updated category. Id: %d.", category.Id)
}

func (repository *categoryRepository) DeleteCategory(categoryId int64) {
	log.Printf("Deleting category. Id: %d", categoryId)
	result, err := repository.stmt.DeleteById.Exec(categoryId)
	if err != nil {
		log.Println(err)
		exception.PanicBadRequest(enum.CategoryNotDelete)
	}
	if result != nil {

	}
	log.Printf("Deleted category. Id: %d", categoryId)
}

func loopThroughCategoryRows(rows *sql.Rows) []entity.Category {
	var categories []entity.Category
	for rows.Next() {
		categories = append(categories, *rowToCategory(rows))
	}
	return categories
}

func rowToCategory(row database.Row) *entity.Category {
	var category entity.Category
	err := row.Scan(&category.Id, &category.CategoryName)
	if err != nil {
		log.Println(err)
		exception.PanicNotFound(enum.CategoryNotFound)
	}
	return &category
}

func checkUniqueException(message string, err error) {
	if err == nil {
		return
	}
	log.Println(err)
	switch err.(type) {
	case *pq.Error:
		if err.(*pq.Error).Code == "23505" {
			message = enum.CategoryAlreadyExists
		}
	}
	exception.PanicBadRequest(message)
}

func closeRows(rows *sql.Rows, message string) {
	err := rows.Close()
	if err != nil {
		log.Println(err)
		exception.PanicNotFound(message)
	}
}
