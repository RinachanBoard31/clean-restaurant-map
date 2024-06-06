package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type TestStruct struct {
	Message string
}

func TestFunc(c echo.Context) error {
	testMessage := TestStruct{Message: "Hello"}
	return c.JSON(http.StatusOK, testMessage)
}

func ActivateRouter() {
	e := echo.New()

	e.GET("/", TestFunc)

	e.Logger.Fatal(e.Start(":8080"))
}
