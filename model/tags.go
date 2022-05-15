package model

import "gorm.io/gorm"

type Tags struct {
	gorm.Model
	Popular   string `json:"popular"`
	Pops      string `json:"pops"`
	CommentId uint `gorm:"foreignKey :ID; reference:CommentId" json:"comment_id"`

}

type PaginateTagsResponse struct {
	Resp   []Tags     `json:"resp"`
	Paging Pagination `json:"paging"`
}
