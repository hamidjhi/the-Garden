package model

import "gorm.io/gorm"

type Garden struct {
	gorm.Model
	Name string `json:"name"`
	TreeId uint `gorm:"ForeignKey:TreeId" json:"tree_id"`
}

type PaginateGardenResponse struct {
	Resp []Garden `json:"resp"`
	Paging Pagination `json:"paging"`
}

