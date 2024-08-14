package models

import "gorm.io/gorm"

type Shipping struct {
	gorm.Model
	OrderID uint `json:"order_id"`
	ShippingAddress string `json:"shipping_address"`
	ShippingMethod string `json:"shipping_method"`
	TrackingNumber string `json:"tracking_number"`
	ShippedDate string `json:"shipped_date"`
	DeliveredDate string `json:"delivered_date"`
	ShippingStatus string `json:"shipping_status"`
}

