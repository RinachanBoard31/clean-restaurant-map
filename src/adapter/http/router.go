package router

import (
	"github.com/labstack/echo/v4"
)

func TestFunc() { 
	return {"value": "Hello"}
}

func ActivateRouter() {
	e := echo.New()

	e.GET("/", TestFunc)

	e.Logger.Fatal(e.Start(":8080"))
}