package controller

import (
	"chemex/logic"
	"chemex/model"
	"chemex/utils"
	"github.com/labstack/echo"
	"net/http"
)

func showTrees(c echo.Context)(err error)  {
	var TreeId, GardenId string
	var date model.Date

	paging, err := getPagePerPage(c, err)

	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	if date.FromDate = c.QueryParam("from_date"); date.FromDate == "" {
		date.FromDate = "2006-01-01"
	}

	if date.ToDate = c.QueryParam("from_date"); date.ToDate == "" {
		date.ToDate = "2050-01-01"
	}

	if TreeId = c.QueryParam("treeId") ; TreeId == ""{
		TreeId = ""
	}
	if GardenId = c.QueryParam("gardenId"); GardenId == ""{
		GardenId = ""
	}

	resp, err:= logic.ShowTrees(date, TreeId, GardenId, paging)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK,resp)
}

func showTreesByQr(c echo.Context)(err error)  {

	paging, err := getPagePerPage(c, err)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	qr := c.QueryParam("qr")
    if qr == "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp, err := logic.ShowTreesByQr(qr, paging)
	if err != nil {
		return c.JSON(http.StatusBadGateway, err)
	}
	return c.JSON(http.StatusOK, resp)
}

func createTree(c echo.Context)(err error)  {
   var userId string
      userId = c.QueryParam("userId")

	if userId == "" {
		return c.JSON(http.StatusBadRequest,err.Error()+"you must be enter your id to add tree")
	}
	n := new(model.Tree)
     qr := utils.GenerateQr()
	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	 err = logic.CreateTree(n, userId, qr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error()+"its seems the tree created before")
	}
	return c.JSON(http.StatusCreated,"tree created successfully")
}

func updateTree(c echo.Context)(err error)  {

      str := c.QueryParam("id")
      if str == ""{
      	return c.JSON(http.StatusBadRequest, err)
	  }
	n := new(model.Tree)

	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	 err = logic.UpdateTree(n,str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,"tree updated successfully")

}

func deleteTree(c echo.Context)(err error) {
	var id string

	if id = c.QueryParam("id"); id == "" {
		return c.JSON(http.StatusBadRequest,err)
	}
	err = logic.DeleteTree(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,"tree deleted successfully!!")
}

