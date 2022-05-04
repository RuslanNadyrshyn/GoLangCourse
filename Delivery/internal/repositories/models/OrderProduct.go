package models

type OrderProduct struct {
	Id        int64   `json:"id,omitempty"`
	OrderId   int64   `json:"orderId,order_id,omitempty"`
	ProductId int64   `json:"productId,product_id,omitempty"`
	Count     int64   `json:"count,counter,omitempty"`
	Price     float64 `json:"price,omitempty"`
}
