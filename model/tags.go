package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
    Tags string `json:"tags"`
	CommentId uint `gorm:"foreignKey : TagId" json:"comment_id"`
	Color string `json:"color"`
}

type PaginateTagsResponse struct {
	Resp   []Tags     `json:"resp"`
	Paging Pagination `json:"paging"`
}
