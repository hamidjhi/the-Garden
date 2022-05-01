package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
)

func ShowGardens(date model.Date, gardenId string, paginate *model.Paginate)( *model.PaginateGardenResponse, error)  {
	var all model.PaginateGardenResponse
	if gardenId != "" {
		response := MySQL.Model(&model.Garden{}).
			Where(model.Garden{GardenId: gardenId}).
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
	response := MySQL.Model(&model.Garden{}).
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

func CreateGarden(garden *model.Garden)(resp *model.Garden, err error)  {


	response := MySQL.Model(model.Garden{}).FirstOrCreate(&garden)
	if response.Error != nil{
		return nil, response.Error
	}
	return resp, nil
}

func UpdateGarden(garden *model.Garden, id string)(resp *model.Garden, err error)  {
	response := MySQL.Model(model.Garden{}).Where(model.Garden{GardenId: id}).Updates(&garden)
	if response.Error != nil{
		return nil, response.Error
	}
	return resp, nil
}

func DeleteGarden(gardenId string)(resp *model.Garden, err error)  {
	response := MySQL.Model(model.Garden{}).Delete(&gardenId)
	if response.Error != nil {
		return nil, response.Error
	}
	return resp, nil
}

