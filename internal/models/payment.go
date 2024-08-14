package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	OrderID       uint    `json:"order_id"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"payment_method"`
	PaymentStatus string  `json:"payment_status"`
}
