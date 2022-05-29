package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
	gorm2 "gorm.io/gorm"
	"log"
	"strconv"
)

func ShowComments(date model.Date, commentId, treeId, tagId string, paginate *model.Paginate) (*model.CommentResponsePaginate, error) {
	var all model.CommentResponsePaginate
	if commentId != "" {
		v, err := strconv.Atoi(commentId)
		if err != nil {
			log.Println("cannot convert string to int in show Comments")
		}
		id := uint(v)

		response := MySQL.Model(&model.Comment{}).
			Where(&model.Comment{Model: gorm2.Model{ID: id}}).
			Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate).Preload("tags")
		if response != nil {
			return nil, response.Error
		}
		p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
		p.SetPage(paginate.Page)

		if err := p.Results(&all.Resp); err != nil {
			return nil, err
		}

		all.Paging.Next, _ = p.NextPage() // 8
		all.Paging.Perv, _ = p.PrevPage() // 6
		all.Paging.Current = p.Page()     // 7
		totalPage := p.Nums()             // 7
		all.Paging.Total = totalPage
		return &all, nil
	}

	if treeId != "" && tagId != ""{
		v, err := strconv.Atoi(treeId)
		v1, err := strconv.Atoi(tagId)
		if err != nil {
			log.Println("cannot convert string to int from show Comments")
		}
		treeeid := uint(v)
		taggId  := uint(v1)
		response := MySQL.Table("comments").
			Where("tree_id = ?", treeeid).Where("tag_id = ?", taggId)
		if response.Error != nil {
			return nil, response.Error
		}
		p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
		p.SetPage(paginate.Page)

		if err := p.Results(&all.Resp); err != nil {
			return nil, err
		}

		all.Paging.Next, _ = p.NextPage() // 8
		all.Paging.Perv, _ = p.PrevPage() // 6
		all.Paging.Current = p.Page()     // 7
		totalPage := p.Nums()             // 7
		all.Paging.Total = totalPage
		return &all, nil
	}
	response := MySQL.Model(&model.Comment{}).
		Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
	if response.Error != nil {
		return nil, response.Error
	}
	p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
	p.SetPage(paginate.Page)

	if err := p.Results(&all.Resp); err != nil {
		return nil, err
	}

	all.Paging.Next, _ = p.NextPage() // 8
	all.Paging.Perv, _ = p.PrevPage() // 6
	all.Paging.Current = p.Page()     // 7
	totalPage := p.Nums()             // 7
	all.Paging.Total = totalPage
	return &all, nil

}

func CreateComment(comment *model.Comment) (err error) {

	response := MySQL.Model(model.Comment{}).Create(&comment)
	if response.Error != nil {
		return err
	}
	return nil
}

func UpdateComment(comment *model.Comment, id string) (err error) {
	v, err := strconv.Atoi(id)
	if err != nil {
		log.Println("cannot convert")
	}
	ids := uint(v)
	response := MySQL.Model(model.Comment{}).Where(&model.Comment{Model: gorm2.Model{ID: ids}}).Updates(&comment)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func DeleteComment(commentId string) (err error) {
	v, err := strconv.Atoi(commentId)
	if err != nil {
		log.Println("cannot convert")
	}
	id := uint(v)
	response := MySQL.Table("comments").Where("id = ?", id).Delete(&id)
	if response.Error != nil {
		return  response.Error
	}
	return  nil
}
