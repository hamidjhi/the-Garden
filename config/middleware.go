package config

import (
	"chemex/auth"
	"chemex/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func RegisterMiddlewares(e *echo.Echo) {

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(ServiceStatus)
	e.HTTPErrorHandler = customHTTPErrorHandler

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

}

//ServiceStatus check service status middleware
func ServiceStatus(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if ServiceRestarting {
			return utils.GetNotFoundError("سرویس در حال بارگزاری مجدد است، مدتی دیگر تلاش نمایید.")
		}

		return next(c)
	}
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var jsonBody interface{}
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		jsonBody = he.Message
	}
	if be, ok := err.(utils.HTTPError); ok {
		c.JSON(be.StatusCode, be)
		return
	}
	if be, ok := err.(auth.HTTPError); ok {
		c.JSON(be.StatusCode, be)
		return
	}
	if err = c.JSON(code, jsonBody); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}
