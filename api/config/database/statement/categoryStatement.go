package statement

import (
	"database/sql"
	"go-restaurant-api/api/config/database"
	"log"
)

const (
	getAllCategoriesQuery = "" +
		"SELECT id, category_name " +
		"FROM categories"
	getCategoryQuery = "" +
		"SELECT id, category_name " +
		"FROM categories " +
		"WHERE id = $1"
	createCategoryQuery = "" +
		"INSERT INTO categories(category_name) " +
		"VALUES ($1) " +
		"RETURNING id"
	updateCategoryQuery = "" +
		"UPDATE categories " +
		"SET category_name=$1 " +
		"WHERE id = $2"
	deleteCategoryQuery = "" +
		"DELETE FROM categories " +
		"WHERE id = $1"
)

var CategoryStmt CategoryStatement

func init() {
	CategoryStmt = CategoryStatement{}
	CategoryStmt.new(database.DB)
}

type CategoryStatement struct {
	FindAll    *sql.Stmt
	FindById   *sql.Stmt
	Create     *sql.Stmt
	Update     *sql.Stmt
	DeleteById *sql.Stmt
}

func (categoryStmt *CategoryStatement) new(db *sql.DB) {
	categoryStmt.FindAll = categoryStmt.prepare(db, getAllCategoriesQuery)
	categoryStmt.FindById = categoryStmt.prepare(db, getCategoryQuery)
	categoryStmt.Create = categoryStmt.prepare(db, createCategoryQuery)
	categoryStmt.Update = categoryStmt.prepare(db, updateCategoryQuery)
	categoryStmt.DeleteById = categoryStmt.prepare(db, deleteCategoryQuery)
}

func (categoryStmt *CategoryStatement) prepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	return stmt
}
