package statement

import (
	"database/sql"
	"go-restaurant-api/api/config/database"
	"log"
)

const (
	getAllProductsQuery = "" +
		"SELECT products.id, product_name, description, price, created_at, category_id, categories.category_name " +
		"FROM products " +
		"LEFT JOIN categories ON category_id = categories.id"
	getProductQuery    = getAllProductsQuery + " WHERE products.id = $1"
	createProductQuery = "" +
		"INSERT INTO products(product_name, description, price, category_id) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id, created_at"
	updateProductQuery = "" +
		"UPDATE products " +
		"SET product_name = $1, description = $2, price = $3, category_id = $4 " +
		"WHERE id = $5"
	deleteProductQuery = "" +
		"DELETE FROM products " +
		"WHERE id = $1"
)

var ProductStmt ProductStatement

func init() {
	ProductStmt = ProductStatement{}
	ProductStmt.new(database.DB)
}

type ProductStatement struct {
	FindAll    *sql.Stmt
	FindById   *sql.Stmt
	Create     *sql.Stmt
	Update     *sql.Stmt
	DeleteById *sql.Stmt
}

func (productStmt *ProductStatement) new(db *sql.DB) {
	productStmt.FindAll = productStmt.prepare(db, getAllProductsQuery)
	productStmt.FindById = productStmt.prepare(db, getProductQuery)
	productStmt.Create = productStmt.prepare(db, createProductQuery)
	productStmt.Update = productStmt.prepare(db, updateProductQuery)
	productStmt.DeleteById = productStmt.prepare(db, deleteProductQuery)
}

func (productStmt *ProductStatement) prepare(db *sql.DB, query string) *sql.Stmt {
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}
	return stmt
}
