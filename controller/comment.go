package controller

import (
	"chemex/logic"
	"chemex/model"


	"github.com/labstack/echo"


	"net/http"

	)

func showComments(c echo.Context)(err error)  {
	var CommentId string
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

	if CommentId = c.QueryParam("commentId") ; CommentId == ""{
		CommentId = ""
	}

	resp, err:= logic.ShowComments(date, CommentId, paging)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK,resp)
}

func createComment(c echo.Context)(err error)  {

	n := new(model.Comment)

	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	resp, err := logic.CreateComment(n)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,resp.CreatedAt)
}

func updateComment(c echo.Context)(err error)  {
     Id := c.QueryParam("CommentId")
	n := new(model.Comment)

	err = c.Bind(&n)
	if err != nil {
		return c.JSON(http.StatusBadRequest,err.Error())
	}

	resp, err := logic.UpdateComment(n, Id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated,resp.UpdatedAt)

}

func deleteComment(c echo.Context)(err error) {
	var CommentId string

	if CommentId = c.QueryParam("commentId"); CommentId == "" {
		return c.JSON(http.StatusBadRequest, err)
	}
	resp ,err := logic.DeleteComment(CommentId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError,err.Error())
	}
	return c.JSON(http.StatusOK,resp.DeletedAt)
}

