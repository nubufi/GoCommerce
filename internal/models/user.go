package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	UserID    string `json:"user_id" gorm:"unique"`
	Role      string `json:"role"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
