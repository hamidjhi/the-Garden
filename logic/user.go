package logic

import (
	"chemex/db"
	"chemex/model"
	"log"
	"strconv"
)

func CreateUser(c *model.User)(err error)  {
	  err = db.CreateUser(c)
	if err != nil {
		log.Println("something went wrong to create user!!1")
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