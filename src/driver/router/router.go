package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/adapter/gateway"
	"clean-storemap-api/src/adapter/presenter"
	"clean-storemap-api/src/driver/db"
	"clean-storemap-api/src/usecase/interactor"

	"github.com/labstack/echo/v4"
)

func ActivateRouter() {
	e := echo.New()

	storeDriver := NewDriverFactory()
	storeOutputPort := NewOutputFactory()
	storeInputPort := NewInputFactory()
	storeRepository := NewRepositoryFactory()
	controller := controller.NewStoreController(storeDriver, storeOutputPort, storeInputPort, storeRepository)

	e.GET("/", controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}

func NewDriverFactory() controller.DriverFactory {
	return &db.DbStoreDriver{}
}

func NewOutputFactory() controller.OutputFactory {
	return presenter.NewStoreOutputPort
}

func NewInputFactory() controller.InputFactory {
	return interactor.NewStoreInputPort
}

func NewRepositoryFactory() controller.RepositoryFactory {
	return gateway.NewStoreRepository
}
