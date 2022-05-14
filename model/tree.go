package model

import (
	"gorm.io/gorm"
)

type Tree struct {
	gorm.Model
	FullName    string `json:"full_name"`
	Age         string `json:"age"`
	DateOfBirth string `json:"date_of_birth"`
	Type        string  `json:"type"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Qr          string  `json:"qr"`
	Length      float64 `json:"length"`
	GardenId uint `json:"garden_id"`
	CommentId   uint `json:"comment_id"`
	Garden []Garden `gorm:"foreignKey:ID; reference:GardenId" `
	Comment []Comment `json:"comment"`

}

type PaginateTreeResponse struct {
	Resp   []Tree     `json:"resp"`
	Paging Pagination `json:"paging"`
}


type TreeResult struct {
	gorm.Model
	FullName    string `json:"full_name"`
	Age         string `json:"age"`
	DateOfBirth string `json:"date_of_birth"`
	Type        string  `json:"type"`
	Lat         float64 `json:"lat"`
	Long        float64 `json:"long"`
	Qr          string  `json:"qr"`
	Length      float64 `json:"length"`
	Comment []Comment `gorm:"ForeignKey:tree_id;AssociationForeignKey:id"`

}

type Result struct {
	ID uint `json:"id"`
	FullName string `json:"full_name"`
	Lat float64 `json:"lat"`
	Text string `json:"text"`
	Name string `json:"name"`
}

