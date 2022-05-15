package logic

import (
	"chemex/db"
	"chemex/model"
	"log"
)

func ShowComments(date model.Date, commentId string, page *model.Paginate, )(resp *model.CommentResponsePaginate,err error) {

	resp, err = db.ShowComments(date, commentId, page)
	if err != nil {
		log.Println("something wrong on show comments, please check!")
	}

	return resp, nil
}

func CreateComment(comment *model.Comment)(err error)  {
	err = db.CreateComment(comment)
	if err != nil {
		log.Println("something wrong on create comment, please check!")
	}
	return  nil
}

func UpdateComment(comment *model.Comment, id string)(err error)  {
	 err = db.UpdateComment(comment, id)
	if err != nil {
		log.Println("something wrong on updating comment, please check!")
	}
	return  nil
}

func DeleteComment(CommentId string)(err error) {

	 err = db.DeleteComment(CommentId)
	if err != nil {
		log.Println("something wrong on deleting comment, please check!!  ")
	}
	return nil
}

