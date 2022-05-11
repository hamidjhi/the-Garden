package model

import (
	"gorm.io/gorm"
	"time"
)

type Tree struct {
	gorm.Model
	FullName string `json:"full_name"`
	Age string `json:"age"`
	DateOfBirth time.Time
	Type string `json:"type"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
	Qr string `json:"qr"`
	Length float64 `json:"length"`
}

type PaginateTreeResponse struct {
	Resp []Tree `json:"resp"`
	Paging Pagination `json:"paging"`
}
