package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	IsActive    bool   `json:"is_active"`
	Garden      []Garden `gorm:"foreignKey:ID" json:"garden"`
}

type PaginateUserResponse struct {
	User  []User `json:"user"`
	Paging Pagination `json:"paging"`
}