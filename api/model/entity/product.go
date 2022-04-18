package entity

import "time"

type Product struct {
	Id          int64
	ProductName string
	Description string
	Price       float64
	CreatedAt   time.Time
	Category    Category
}
