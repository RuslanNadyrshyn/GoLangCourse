package models

import "time"

type Product struct {
	Id 			int
	Date 		time.Time
	SupplierId 	int
	Price		float32
}
