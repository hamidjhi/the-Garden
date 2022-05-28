package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Popular string `json:"popular"`
	Pops    string `json:"pops"`
	Pic     string `json:"pic"`
}

type PaginateTagsResponse struct {
	Resp   []Tags     `json:"resp"`
	Paging Pagination `json:"paging"`
}
