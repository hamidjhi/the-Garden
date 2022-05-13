package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Popular   string `json:"popular"`
	Pops      string `json:"pops"`
	CommentId uint   `gorm:"ForeignKey:comment_id" json:"comment_id"`
}

type PaginateTagsResponse struct {
	Resp   []Tags     `json:"resp"`
	Paging Pagination `json:"paging"`
}
