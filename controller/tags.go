package controller

import (
	"chemex/logic"
	"chemex/model"
	"github.com/labstack/echo"
	"net/http"
)

func showTags(c echo.Context) (err error) {
	var TagId string
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

	if TagId = c.QueryParam("tag_id"); TagId == "" {
		TagId = ""
	}

	resp, err := logic.ShowTags(date, TagId, paging)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, resp)

}

func addTag(c echo.Context) (err error) {

	n := new(model.Tags)

	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = logic.AddTag(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, "tag created")

}

func updateTag(c echo.Context) (err error) {

	id := c.QueryParam("tag_id")
	n := new(model.Tags)

	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = logic.UpdateTag(n, id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, "tag updated")

}

func deleteTag(c echo.Context) (err error) {

	var TagId string

	if TagId = c.QueryParam("tag_id"); TagId == "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	err = logic.DeleteTag(TagId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "tag deleted")

}
