package models

import "time"

type Product struct {
	Id         int
	SupplierId int
	Price      float32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
