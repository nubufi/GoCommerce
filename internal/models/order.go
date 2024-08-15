package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          uint    `json:"user_id"`
	TotalPrice      float64 `json:"total_price"`
	OrderStatus     string  `json:"order_status"`
	PaymentMethod   string  `json:"payment_method"`
	PaymentStatus   string  `json:"payment_status"`
	OrderDate       string  `json:"order_date"`
	ShippingAddress string  `json:"shipping_address"`
	OrderItems      []OrderItem `json:"order_items"`
}

type OrderItem struct {
	gorm.Model
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
