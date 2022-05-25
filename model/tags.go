package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Popular   string `json:"popular"`
	Pops      string `json:"pops"`
	Comment []Comment `gorm:"foreignKey: ID" json:"comment"`
	Pic string `json:"pic"`
}

type PaginateTagsResponse struct {
	Resp   []Tags     `json:"resp"`
	Paging Pagination `json:"paging"`
}
