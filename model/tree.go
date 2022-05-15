package model

import (
	"gorm.io/gorm"
	"net/url"
)

type Tree struct {
	gorm.Model
	FullName    string  `json:"full_name"`
	Age         string  `json:"age"`
	DateOfBirth string  `json:"date_of_birth"`
	Type        string  `json:"type"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Qr          string  `json:"qr"`
	Length      float64 `json:"length"`
	GardenId    uint    `json:"garden_id"`
	Pic         url.URL `json:"pic"`
	CommentId   uint    `json:"comment_id"`
	Description string  `json:"description"`
}

type PaginateTreeResponse struct {
	Resp   []Tree     `json:"resp"`
	Paging Pagination `json:"paging"`
}




type Result struct {
	ID uint `json:"id"`
	FullName string `json:"full_name"`
	Lat float64 `json:"lat"`
	Text string `json:"text"`
	Name string `json:"name"`
}

