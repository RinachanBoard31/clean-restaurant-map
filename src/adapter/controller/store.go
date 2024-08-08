package controller

import (
	"clean-storemap-api/src/adapter/gateway"
	"clean-storemap-api/src/usecase/port"

	"github.com/labstack/echo/v4"
)

type storeForController struct {
	Id   int
	Name string
}

type StoreJson struct {
	ResponseCode int
	Message      string
	Stores       []storeForController
}

type StoreI interface {
	GetStores(c echo.Context) error
}

type StoreOutputFactory func(echo.Context) port.StoreOutputPort
type StoreInputFactory func(port.StoreRepository, port.StoreOutputPort) port.StoreInputPort
type StoreRepositoryFactory func(gateway.StoreDriver) port.StoreRepository
type StoreDriverFactory gateway.StoreDriver

type StoreController struct {
	storeDriverFactory     StoreDriverFactory
	storeOutputFactory     StoreOutputFactory
	storeInputFactory      StoreInputFactory
	storeRepositoryFactory StoreRepositoryFactory
}

func NewStoreController(storeDriverFactory StoreDriverFactory, storeOutputFactory StoreOutputFactory, storeInputFactory StoreInputFactory, storeRepositoryFactory StoreRepositoryFactory) StoreI {
	return &StoreController{
		storeDriverFactory:     storeDriverFactory,
		storeOutputFactory:     storeOutputFactory,
		storeInputFactory:      storeInputFactory,
		storeRepositoryFactory: storeRepositoryFactory,
	}
}

func (sc *StoreController) GetStores(c echo.Context) error {
	return sc.newStoreInputPort(c).GetStores()
}

/* ここでpresenterにecho.Contextを渡している！起爆！！！（遅延） */
/* これによって、presenterのinterface(outputport)にecho.Contextを書かなくて良くなる */
func (sc *StoreController) newStoreInputPort(c echo.Context) port.StoreInputPort {
	storeOutputPort := sc.storeOutputFactory(c)
	storeDriver := sc.storeDriverFactory
	storeRepository := sc.storeRepositoryFactory(storeDriver)
	return sc.storeInputFactory(storeRepository, storeOutputPort)
}
