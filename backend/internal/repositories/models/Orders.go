package models

type Order struct {
	Id        int64   `json:"id"`
	Price     float64 `json:"price"`
	UserId    int64   `json:"userId"`
	Address   string  `json:"address"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
