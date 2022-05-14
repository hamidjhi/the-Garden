package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Popular   string `json:"popular"`
	Pops      string `json:"pops"`
	CommentId uint `json:"comment_id"`
	Comment Comment `gorm:"foreignKey :ID; reference:CommentId" json:"comment"`
}

type PaginateTagsResponse struct {
	Resp   []Tags     `json:"resp"`
	Paging Pagination `json:"paging"`
}
