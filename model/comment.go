package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	TreeId uint  `gorm:"ForeignKey:tree_id" json:"comment_id"`
}

type CommentResponsePaginate struct {
	Resp []Comment `json:"resp"`
	Paging Pagination `json:"paging"`
}