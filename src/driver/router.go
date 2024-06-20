package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/usecase/interactor"

	"github.com/labstack/echo/v4"
)

func ActivateRouter() {
	e := echo.New()

	// コンストラクタを呼び出していないのでエラーになる。
	// RepositoryとOutputの実装が完了しないとエラーは解消しない。
	inputPort := interactor.NewStoreInputPort()
	controller := controller.NewStoreController(inputPort)

	e.GET("/", controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}
