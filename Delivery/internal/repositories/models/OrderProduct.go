package models

type OrderProduct struct {
	Id        int     `json:"id,omitempty"`
	OrderId   int     `json:"orderId,order_id,omitempty"`
	ProductId int     `json:"productId,product_id,omitempty"`
	Count     int     `json:"count,counter,omitempty"`
	Price     float64 `json:"price,omitempty"`
}
