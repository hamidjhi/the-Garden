package model

import (
	"gorm.io/gorm"
)

type Garden struct {
	gorm.Model
	Name          string  `json:"name"`
	Lat           float64 `json:"lat"`
	Long          float64 `json:"long"`
	Pic           string  `json:"pic"`
	NumberOfTrees uint    `json:"number_of_trees"`
	UserId        uint    `json:"user_id"`
	Tree          []Tree  `gorm:"foreignKey:ID" json:"tree"`
}

type PaginateGardenResponse struct {
	Resp []Garden `json:"resp"`
	Paging Pagination `json:"paging"`
}

