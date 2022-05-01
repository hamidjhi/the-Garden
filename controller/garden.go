package controller

import (
	"chemex/logic"
	"chemex/model"
	"github.com/labstack/echo"
	"net/http"
)

func showGardens(c echo.Context)(err error)  {
    var gardenId string
   var date model.Date

    paging, err := getPagePerPage(c, err)

	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	if date.FromDate = c.QueryParam("from_date"); date.FromDate == "" {
		date.FromDate = "2006-01-01"
	}

	if date.FromDate = c.QueryParam("from_date"); date.FromDate == "" {
		date.FromDate = "2050-01-01"
	}

	if gardenId = c.QueryParam("gardenId") ; gardenId == ""{
		gardenId = ""
	}

	resp, err:= logic.ShowGardens(date, gardenId, paging)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK,resp)
}

func createGarden(c echo.Context)(err error)  {

	n := new(model.Garden)

	 err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	 resp, err := logic.CreateGarden(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,resp.CreatedAt)
}

func updateGarden(c echo.Context)(err error)  {
     id := c.QueryParam("GardenId")
	n := new(model.Garden)

	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	resp, err := logic.UpdateGarden(n, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,resp.UpdatedAt)

}

func deleteGarden(c echo.Context)(err error) {
	var gardenId string

	if gardenId = c.QueryParam("gardenId"); gardenId == "" {
		return c.JSON(http.StatusBadRequest,err)
	}
	 resp ,err := logic.DeleteGarden(gardenId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
      return c.JSON(http.StatusOK,resp.DeletedAt)
}

