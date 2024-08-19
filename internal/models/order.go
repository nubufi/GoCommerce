package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID          string      `json:"user_id" binding:"required"`
	TotalPrice      float64     `json:"total_price" binding:"required"`
	PaymentMethod   string      `json:"payment_method" binding:"required"`
	ShippingAddress string      `json:"shipping_address" binding:"required"`
	OrderItems      []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID    uint    `json:"order_id"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
}
