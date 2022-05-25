package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	TreeId uint `json:"tree_id"`
	Pic string `json:"pic"`
}

type CommentResponsePaginate struct {
	Resp []Comment `json:"resp"`
	Paging Pagination `json:"paging"`
}