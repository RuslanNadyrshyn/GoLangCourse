package models

type Product struct {
	Id          int      `json:"id"`
	Name        string   `json:"name"`
	MenuId      int      `json:"menuId"`
	Price       float64  `json:"price"`
	Image       string   `json:"image"`
	Type        string   `json:"type"`
	Ingredients []string `json:"ingredients"`
}
