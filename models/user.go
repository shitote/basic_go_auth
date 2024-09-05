package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"eamil"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
