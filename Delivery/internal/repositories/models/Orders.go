package models

type Order struct {
	Id      int     `json:"id"`
	Price   float64 `json:"price"`
	UserId  int     `json:"userId"`
	Address string  `json:"address"`
}
