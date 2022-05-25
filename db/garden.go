package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func ShowGardens(date model.Date, gardenId string, UserId string, paginate *model.Paginate) (*model.PaginateGardenResponse, error) {
	var all model.PaginateGardenResponse
	if gardenId != "" {

		v, err := strconv.Atoi(gardenId)
		if err != nil {
			log.Println("cannot convert string to int")
		}
		id := uint(v)
		response := MySQL.Model(&model.Garden{}).
			Where(&model.Garden{Model: gorm.Model{ID: id}}).
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

	if UserId != "" {

		v, err := strconv.Atoi(UserId)
		if err != nil {
			log.Println("cannot convert string to int")
		}
		id := uint(v)
		response := MySQL.Model(&model.Garden{}).
			Where(&model.Garden{Model: gorm.Model{ID: id}}).
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

	response := MySQL.Model(&model.Garden{}).
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

func ShowGardenByNumber(number string)([]*model.Garden, error)  {
	var resp []*model.Garden
	var err error

	res := MySQL.Model(&model.Garden{}).Where(model.Garden{AdminNumber: number}).Find(&resp)
	if res == nil{
		return nil, err
	}
	return resp, nil
}


func CreateGarden(garden *model.Garden) (err error) {

	response := MySQL.Model(model.Garden{}).Create(&garden)
	if response.Error != nil {
		return err
	}
	return nil
}

func UpdateGarden(garden *model.Garden, str string) (err error) {

	v, err := strconv.Atoi(str)
	id := uint(v)
	response := MySQL.Model(model.Garden{}).Where(&model.Garden{Model: gorm.Model{ID: id}}).Updates(&garden)
	if response.Error != nil {
		return err
	}
	return nil
}

func DeleteGarden(gardenId string) (err error) {
	v, err := strconv.Atoi(gardenId)
	id := uint(v)
	response := MySQL.Table("trees").Where("id = ?", id).Delete(&id)
	if response.Error != nil {
		return err
	}
	return nil
}
