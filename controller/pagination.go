package controller

import (
	"chemex/model"
	"errors"
	"github.com/labstack/echo"
	"strconv"
)

func getPagePerPage(ctx echo.Context, err error) (*model.Paginate, error) {
	paging := new(model.Paginate)
	if ctx.QueryParam("page") != "" {
		paging.Page, err = strconv.Atoi(ctx.QueryParam("page"))
		if err != nil {
			return nil, errors.New(err.Error())
		}
	} else {
		paging.Page = 1
	}
	if ctx.QueryParam("per_page") != "" {
		paging.PerPage, err = strconv.Atoi(ctx.QueryParam("per_page"))
		if err != nil {
			return nil, errors.New(err.Error())
		}
	} else {
		paging.PerPage = 10
	}
	return paging, nil
}

