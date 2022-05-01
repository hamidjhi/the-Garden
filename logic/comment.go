package logic

import (
	"chemex/db"
	"chemex/model"
	"log"
)

func ShowComments(date model.Date, commentId string, page *model.Paginate, )(resp *model.CommentResponsePaginate,err error) {

	resp, err = db.ShowComments(date, commentId, page)
	if err != nil {
		log.Println("we have an err to getting response from comment table")
	}

	return resp, nil
}

func CreateComment(comment *model.Comment)(resp *model.Comment, err error)  {
	resp, err = db.CreateComment(comment)
	if err != nil {
		log.Println("cannot add anything to garden table ")
	}
	return resp, nil
}

func UpdateComment(comment *model.Comment, id string)(resp *model.Comment, err error)  {
	resp, err = db.UpdateComment(comment, id)
	if err != nil {
		log.Println("cannot update comment table")
	}
	return resp, nil
}

func DeleteComment(CommentId string)(resp *model.Comment, err error) {

	resp, err = db.DeleteComment(CommentId)
	if err != nil {
		log.Println("cannot ")
	}
	return resp, nil
}

