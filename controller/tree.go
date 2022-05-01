package controller

import (
	"chemex/logic"
	"chemex/model"
	"chemex/utils"
	"github.com/labstack/echo"
	"net/http"
)

func showTrees(c echo.Context)(err error)  {
	var TreeId string
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

	resp, err:= logic.ShowTrees(date, TreeId, paging)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK,resp)
}

func showTreesByQr(c echo.Context)  {






}







func createTree(c echo.Context)(err error)  {

	n := new(model.Tree)
     qr := utils.GenerateQr()
	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	 err = logic.CreateTree(n, qr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error()+"its seems the user created before")
	}
	return c.JSON(http.StatusCreated,"user created successfully")
}

func updateTree(c echo.Context)(err error)  {

      str := c.QueryParam("treeId")
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
	var treeId string

	if treeId = c.QueryParam("treeId"); treeId == "" {
		return c.JSON(http.StatusBadRequest,err)
	}
	resp ,err := logic.DeleteTree(treeId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,resp.DeletedAt)
}

