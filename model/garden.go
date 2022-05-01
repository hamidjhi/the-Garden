package model

import "gorm.io/gorm"

type Garden struct {
	gorm.Model
	Name string `json:"name"`
	GardenId string `json:"garden_id"`
	Lat string `json:"lat"`
	Long string `json:"long"`
}

type PaginateGardenResponse struct {
	Resp []Garden `json:"resp"`
	Paging Pagination `json:"paging"`
}

