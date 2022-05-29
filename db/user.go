package db

import (
	"chemex/model"
	"github.com/vcraescu/go-paginator"
	"github.com/vcraescu/go-paginator/adapter"
	gorm2 "gorm.io/gorm"
	"log"
	"strconv"
)

func GetUser(date model.Date, userId string, paginate *model.Paginate) (*model.PaginateUserResponse, error)  {

	var all model.PaginateUserResponse
	var err error

	if userId != "" {

		v, err := strconv.Atoi(userId)
		if err != nil {
			log.Println("cannot convert")
		}
		id := uint(v)

		response := MySQL.Model(&model.User{}).
			Where(&model.User{Model: gorm2.Model{ID: id}}).Preload("Garden")
		if response.Error != nil {
			return nil, response.Error
		}
		p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
		p.SetPage(paginate.Page)

		if err := p.Results(&all.User); err != nil {
			return nil, err
		}

		all.Paging.Next, _ = p.NextPage() // 8
		all.Paging.Perv, _ = p.PrevPage() // 6
		all.Paging.Current = p.Page()     // 7
		totalPage := p.Nums()             // 7
		all.Paging.Total = totalPage
		return &all, nil
	}
	response := MySQL.Table("users").
		Where("created_at BETWEEN ? AND ?", date.FromDate, date.ToDate)
	if response == nil {
		return nil, err
	}
	p := paginator.New(adapter.NewGORMAdapter(response), paginate.PerPage)
	p.SetPage(paginate.Page)

	if err = p.Results(&all.User); err != nil {
		return nil, err
	}

	all.Paging.Next, _ = p.NextPage() // 8
	all.Paging.Perv, _ = p.PrevPage() // 6
	all.Paging.Current = p.Page()     // 7
	totalPage := p.Nums()             // 7
	all.Paging.Total = totalPage
	return &all, nil

}
func ShowUserByNumber(number string)([]*model.User, error)  {
	var resp []*model.User
	var err error

	res := MySQL.Model(&model.User{}).Where(model.User{PhoneNumber: number}).Find(&resp)
	if res == nil{
		return nil, err
	}
	return resp, nil
}

func CreateUser(c *model.User)(err error)  {

	response := MySQL.Model(model.User{}).Create(&c).Error
	if response.Error != nil {
		return err
	}
	return nil

}

func UpdateUser(userId uint, c *model.User)(err error)  {
	response := MySQL.Model(model.User{}).Where(&model.User{Model: gorm2.Model{ID: userId}}).Updates(&c).Error
	if response.Error != nil {
		return err
	}
	return nil
}