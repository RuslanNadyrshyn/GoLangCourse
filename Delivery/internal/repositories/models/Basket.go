package models

import "time"

type Basket struct {
	Id         int
	UserId     int
	SupplierId int
	Price      float32
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
