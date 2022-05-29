package logic

import (
	"chemex/db"
	"chemex/model"
	"log"
	"strconv"
)

func GetUser(date model.Date, userId string, page *model.Paginate)(resp *model.PaginateUserResponse,err error)  {
	resp, err = db.GetUser(date, userId, page)
	if err != nil {
		log.Println("we have an err to getting response from tree table")
	}

	return resp, nil
}


func ShowUserByNumber(number string)(res []*model.User, err error)  {
	res, err = db.ShowUserByNumber(number)
	if err != nil {
		log.Println("someThing wrong to response data from user")
	}
	return res, nil
}


func CreateUser(c *model.User)(err error)  {
	  err = db.CreateUser(c)
	if err != nil {
		log.Println("something went wrong to create user!!!")
	}
	return nil
}

func UpdateUser(userId string, c *model.User)(err error)  {
	v , err:= strconv.Atoi(userId)
	if err != nil {
		log.Println("cannot convert")
	}
	id := uint(v)
	err = db.UpdateUser(id, c)
	if err != nil {
		log.Println("something went wrong when updating user")
	}
	return nil
}