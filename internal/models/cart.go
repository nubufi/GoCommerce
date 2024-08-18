package models

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model
	UserID    string  `json:"user_id" binding:"required"`
	ProductID uint  `json:"product_id" binding:"required" `
	Quantity  int     `json:"quantity" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
}
