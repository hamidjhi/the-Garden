package model

import (
	"gorm.io/gorm"
)

type Garden struct {
	gorm.Model
	Name string `json:"name"`
	TreeId uint `gorm:"ForeignKey:TreeId" json:"tree_id"`
	Lat float64 `json:"lat"`
	Long float64 `json:"long"`
	AdminNumber   string `json:"admin_number"`
	AdminName string `json:"admin_name"`
	Pic string`json:"pic"`
}

type PaginateGardenResponse struct {
	Resp []Garden `json:"resp"`
	Paging Pagination `json:"paging"`
}

type GardenLocation struct {
	gorm.Model
	Lat1 float64 `json:"lat_1"`
	Lat2 float64 `json:"lat_2"`
	Lat3 float64 `json:"lat_3"`
	Lat4 float64 `json:"lat_4"`
	Long1 float64 `json:"long_1"`
	Long2 float64 `json:"long_2"`
	Long3 float64 `json:"long_3"`
	Long4 float64 `json:"long_4"`
	GardenId uint  `gorm:"ForeignKey:GardenId" json:"garden_id"`
}