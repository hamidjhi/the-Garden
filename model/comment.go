package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	Tag []Tags  `gorm:"foreignKey:ID" json:"tag"`
	Pic string `json:"pic"`
}

type CommentResponsePaginate struct {
	Resp []Comment `json:"resp"`
	Paging Pagination `json:"paging"`
}