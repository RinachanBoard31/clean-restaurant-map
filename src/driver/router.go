package router

import (
	controller "clean-storemap-api/src/adapter/controller"
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

	e.GET("/test", TestFunc)
	// コンストラクタを呼び出していないのでエラーになる
	e.GET("/", controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}
