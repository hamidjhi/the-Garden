package model

import "gorm.io/gorm"

type Tree struct {
	gorm.Model
	Name string `json:"name"`
	TreeId string `json:"tree_id"`
	Qr string `gorm:"default:uuid_generate_v3()" json:"qr"`
	Lat string `json:"lat"`
	Long string `json:"long"`
	Garden []Garden `gorm:"ForeignKey:GardenId" json:"garden"`
}

type PaginateTreeResponse struct {
	Resp []Tree `json:"resp"`
	Paging Pagination `json:"paging"`
}
