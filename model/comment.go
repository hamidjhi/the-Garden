package model

import (
	"gorm.io/gorm"
	"net/url"
)

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	TreeId uint `json:"tree_id"`
	Tree  []Tree  `gorm:"ForeignKey:id;reference:Tree_Id" json:"tree"`
	Pic url.URL `json:"pic"`
}

type CommentResponsePaginate struct {
	Resp []Comment `json:"resp"`
	Paging Pagination `json:"paging"`
}