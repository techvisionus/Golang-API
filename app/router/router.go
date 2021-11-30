package router

import (
	"golang-api/app/controllers"
	"golang-api/app/errors"
	"golang-api/app/middlewares"

	"github.com/labstack/echo/v4"
)

func Init(app *echo.Echo) {
	app.Use(middlewares.Cors())
	app.Use(middlewares.Gzip())
	app.Use(middlewares.Logger())
	app.Use(middlewares.Secure())
	app.Use(middlewares.Recover())
	app.HTTPErrorHandler = errors.HttpErrorHandler

	app.GET("/", controllers.Index())
}
