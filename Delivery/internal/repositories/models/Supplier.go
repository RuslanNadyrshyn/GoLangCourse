package models

import "time"

type Supplier struct {
	Id        string
	Name      string
	Address   string
	Type      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
