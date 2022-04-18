package repository

import (
	"database/sql"
	"go-restaurant-api/api/config/database"
	"go-restaurant-api/api/config/database/statement"
	"go-restaurant-api/api/enum"
	"go-restaurant-api/api/model/entity"
	"go-restaurant-api/api/model/exception"
	"log"
)

var ProductRepo productRepository

type productRepository struct {
	stmt *statement.ProductStatement
}

func init() {
	ProductRepo = productRepository{
		stmt: &statement.ProductStmt,
	}
}

func (repository *productRepository) ListAllProducts() []entity.Product {
	rows, err := repository.stmt.FindAll.Query()
	defer closeRows(rows, enum.ProductNotFound)
	if err != nil {
		log.Panicln(err)
	}
	return loopThroughProductRows(rows)
}

func (repository *productRepository) GetProduct(productId int64) *entity.Product {
	row := repository.stmt.FindById.QueryRow(productId)
	return rowToProduct(row)
}

func (repository *productRepository) CreateProduct(product *entity.Product) {
	log.Printf("Creating product. Name: %s.", product.ProductName)
	err := repository.stmt.Create.
		QueryRow(product.ProductName, product.Description, product.Price, product.Category.Id).
		Scan(&product.Id, &product.CreatedAt)
	if err != nil {
		log.Println(err)
		exception.PanicBadRequest(enum.ProductNotCreated)
	}
	log.Printf("New product. Id: %d, name: %s.", product.Id, product.ProductName)
}

func (repository *productRepository) UpdateProduct(product *entity.Product) {
	log.Printf("Updating category. Id: %d, Name: %s.", product.Id, product.ProductName)
	_, err := repository.stmt.Update.
		Exec(product.ProductName, product.Description, product.Price, product.Category.Id, product.Id)
	if err != nil {
		log.Println(err)
		exception.PanicBadRequest(enum.ProductNotUpdate)
	}
	log.Printf("Updated product. Id: %d, Name: %s.", product.Id, product.ProductName)
}

func (repository *productRepository) DeleteProduct(productId int64) {
	log.Printf("Deleting product. Id: %d", productId)
	_, err := repository.stmt.DeleteById.Exec(productId)
	if err != nil {
		log.Println(err)
		exception.PanicBadRequest(enum.ProductNotDelete)
	}
	log.Printf("Deleted product. Id: %d", productId)
}

func loopThroughProductRows(rows *sql.Rows) []entity.Product {
	var products []entity.Product
	for rows.Next() {
		products = append(products, *rowToProduct(rows))
	}
	return products
}

func rowToProduct(row database.Row) *entity.Product {
	var product entity.Product
	err := row.Scan(
		&product.Id,
		&product.ProductName,
		&product.Description,
		&product.Price,
		&product.CreatedAt,
		&product.Category.Id,
		&product.Category.CategoryName)
	if err != nil {
		log.Println(err)
		exception.PanicNotFound(enum.ProductNotFound)
	}
	return &product
}
