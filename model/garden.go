package model

import (
	"gorm.io/gorm"
	"net/url"
)

type Garden struct {
	gorm.Model
	Name string `json:"name"`
	TreeId uint `gorm:"ForeignKey:TreeId" json:"tree_id"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
	AdminNumber   string `json:"admin_number"`
	AdminName string `json:"admin_name"`
	Pic url.URL`json:"pic"`
}

type PaginateGardenResponse struct {
	Resp []Garden `json:"resp"`
	Paging Pagination `json:"paging"`
}

