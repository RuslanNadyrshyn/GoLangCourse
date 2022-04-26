package models

type Basket struct {
	Id     int     `json:"id"`
	UserId int     `json:"userId"`
	Price  float32 `json:"price"`
}
