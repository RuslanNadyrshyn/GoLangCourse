package models

import "time"

type Orders struct {
	Id        int
	Price     float32
	UserId    int
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
