package main

import (
	"fmt"
	"golang-api/app/router"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/labstack/echo/v4"
	"gopkg.in/tylerb/graceful.v1"
)

func main() {
	app := echo.New()
	router.Init(app)

	app.Server.Addr = ":3333"
	fmt.Println("Server started at ", app.Server.Addr)
	graceful.ListenAndServe(app.Server, 5*time.Second)
}
