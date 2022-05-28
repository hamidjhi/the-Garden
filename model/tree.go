package model

import (
	"gorm.io/gorm"
)

type Tree struct {
	gorm.Model
	FullName    string    `json:"full_name"`
	Age         string    `json:"age"`
	DateOfBirth string    `json:"date_of_birth"`
	Type        string    `json:"type"`
	Lat         float64   `json:"lat"`
	Long        float64   `json:"long"`
	Qr          string    `json:"qr"`
	Length      float64   `json:"length"`
	Pic         string    `json:"pic"`
	Description string    `json:"description"`
	GardenId    uint      `json:"garden_id"`
	Comments    []Comment `gorm:"foreignKey:ID" json:"comments"`
}

type PaginateTreeResponse struct {
	Resp   []Tree     `json:"resp"`
	Paging Pagination `json:"paging"`
}

type TreeResult struct {
	ID uint `json:"id"`
	FullName string `json:"full_name"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
	Description string `json:"description"`
	Text string `json:"text"`
	Name string `json:"name"`
	Type string `json:"type"`
}



