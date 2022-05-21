package controller

import (
	"chemex/logic"
	"chemex/model"
	"github.com/labstack/echo"
	"net/http"
)

func createUser(c echo.Context)(err error)  {

	user := new(model.User)

	       err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	  err = logic.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusBadGateway,err.Error()+"check ur request field or unknown format struct")
	}
	return c.JSON(http.StatusCreated,"user created successfully")

}

func updateUser(c echo.Context)(err error)  {
    str := c.QueryParam("userId")

    if str == ""{
    	return c.JSON(http.StatusBadRequest,err.Error()+"please enter userId!!")
    	}

    	user :=new(model.User)
    	err = c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error()+"we cant bind user model maybe its wrong!")
	}
	  err = logic.UpdateUser(str, user)
	if err != nil {
		return c.JSON(http.StatusBadGateway,err.Error()+"we cant update user check the database!")
	}
	return c.JSON(http.StatusOK,"user created successfully")

}
