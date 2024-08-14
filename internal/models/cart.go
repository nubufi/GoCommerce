package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID uint `json:"user_id"`
}

type CartItem struct {
	gorm.Model
	CartID    uint `json:"cart_id"`
	ProductID uint `json:"product_id"`
	Quantity  int  `json:"quantity"`
}
