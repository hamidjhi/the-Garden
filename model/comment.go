package model

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Text string `json:"text"`
	CommentId string `json:"comment_id"`
	Gard []Garden  `gorm:"ForeignKey:GardenId" json:"gard"`
	Tree []Tree `gorm:"ForeignKey:TreeId" json:"tree"`
}

type CommentResponsePaginate struct {
	Resp []Comment `json:"resp"`
	Paging Pagination `json:"paging"`
}