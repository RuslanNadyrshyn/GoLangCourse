package models

import "time"

type Product struct {
	Id         int
	Name       string
	SupplierId int
	Price      float32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
