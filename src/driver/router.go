package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/usecase/interactor"

	"github.com/labstack/echo/v4"
)

func ActivateRouter() {
	e := echo.New()

	inputPort := interactor.NewStoreInputPort()
	controller := controller.NewStoreController(inputPort)

	// コンストラクタを呼び出していないのでエラーになる
	e.GET("/", controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}
