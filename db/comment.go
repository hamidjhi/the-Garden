package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
)

func ShowComments(date model.Date, commentId string, paginate *model.Paginate)( *model.CommentResponsePaginate, error)  {
	var all model.CommentResponsePaginate
	if commentId != "" {
		response := MySQL.Model(&model.Comment{}).
			Where(model.Comment{CommentId: commentId}).
			Where("created at BETWEEN ? AND ?", date.FromDate, date.ToDate)
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
	response := MySQL.Model(&model.Comment{}).
		Where("created at BETWEEN ? AND ?", date.FromDate, date.ToDate)
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

func CreateComment(commentId *model.Comment)(resp *model.Comment, err error)  {


	response := MySQL.Model(model.Comment{}).FirstOrCreate(&commentId)
	if response.Error != nil{
		return nil, response.Error
	}
	return resp, nil
}

func UpdateComment(comment *model.Comment, id string)(resp *model.Comment, err error)  {
	response := MySQL.Model(model.Comment{}).Where(model.Comment{CommentId: id}).Updates(&comment)
	if response.Error != nil{
		return nil, response.Error
	}
	return resp, nil
}

func DeleteComment(commentId string)(resp *model.Comment, err error)  {
	response := MySQL.Model(model.Comment{}).Delete(&commentId)
	if response.Error != nil {
		return nil, response.Error
	}
	return resp, nil
}


