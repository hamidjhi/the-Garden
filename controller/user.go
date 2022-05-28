package controller

import (
	"chemex/logic"
	"chemex/model"
	"github.com/labstack/echo"
	"net/http"
)

func getUser(c echo.Context)(err error)  {
	var userId string
	var date model.Date

	paging, err := getPagePerPage(c, err)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if date.FromDate = c.QueryParam("from_date"); date.FromDate == "" {
		date.FromDate = "2006-01-01"
	}

	if date.FromDate = c.QueryParam("from_date"); date.FromDate == "" {
		date.FromDate = "2050-01-01"
	}

	if userId = c.QueryParam("user_id"); userId == "" {
		userId = ""
	}

	resp, err := logic.GetUser(date, userId, paging)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}

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
