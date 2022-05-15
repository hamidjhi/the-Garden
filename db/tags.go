package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
	gorm2 "gorm.io/gorm"
	"log"
	"strconv"
)

func ShowTags(date model.Date, tagId string, paginate *model.Paginate) (*model.PaginateTagsResponse, error) {
	var all model.PaginateTagsResponse
	if tagId != "" {

		v, err := strconv.Atoi(tagId)
		if err != nil {
			log.Println("cannot convert string to int")
		}
		id := uint(v)
		response := MySQL.Model(&model.Tags{}).
			Where(&model.Tags{Model: gorm2.Model{ID: id}}).
			Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
		if response == nil {
			return nil, err
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
	response := MySQL.Model(&model.Tags{}).
		Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
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

func AddTag(tags *model.Tags) (err error) {

	response := MySQL.Model(&model.Tags{}).Create(&tags)
	if response.Error != nil {
		return nil
	}
	return nil
}

func UpdateTag(tag *model.Tags, str string) (err error) {

	v, err := strconv.Atoi(str)
	id := uint(v)
	response := MySQL.Model(model.Garden{}).Where(&model.Garden{Model: gorm2.Model{ID: id}}).Updates(&tag)
	if response.Error != nil {
		return response.Error
	}
	return nil
}

func DeleteTag(tagId string) (err error) {
	v, err := strconv.Atoi(tagId)
	id := uint(v)
	response := MySQL.Table("trees").Where("id = ?", id).Delete(&id)
	if response.Error != nil {
		return response.Error
	}
	return nil
}
