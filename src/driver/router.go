package router

import (
	controller "clean-storemap-api/src/adapter/controller"

	"github.com/labstack/echo/v4"
)

func ActivateRouter() {
	e := echo.New()

	controller := controller.NewStoreController()

	// コンストラクタを呼び出していないのでエラーになる
	e.GET("/", controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}
