package db

import (
	"chemex/model"
	gorm2 "gorm.io/gorm"
)

func GetUser(date model.Date, treeId string, paginate *model.Paginate) (*model.PaginateTreeResponse, error)  {

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