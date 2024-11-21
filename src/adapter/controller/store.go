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
	GetNearStores(c echo.Context) error
}

type StoreOutputFactory func(echo.Context) port.StoreOutputPort
type StoreInputFactory func(port.StoreRepository, port.StoreOutputPort) port.StoreInputPort
type StoreRepositoryFactory func(gateway.StoreDriver, gateway.GoogleMapDriver) port.StoreRepository
type StoreDriverFactory gateway.StoreDriver
type GoogleMapDriverFactory gateway.GoogleMapDriver

type StoreController struct {
	storeDriverFactory     StoreDriverFactory
	googleMapDriverFactory GoogleMapDriverFactory
	storeOutputFactory     StoreOutputFactory
	storeInputFactory      StoreInputFactory
	storeRepositoryFactory StoreRepositoryFactory
}

func NewStoreController(
	storeDriverFactory StoreDriverFactory,
	googleMapDriverFactory GoogleMapDriverFactory,
	storeOutputFactory StoreOutputFactory,
	storeInputFactory StoreInputFactory,
	storeRepositoryFactory StoreRepositoryFactory,
) StoreI {
	return &StoreController{
		storeDriverFactory:     storeDriverFactory,
		googleMapDriverFactory: googleMapDriverFactory,
		storeOutputFactory:     storeOutputFactory,
		storeInputFactory:      storeInputFactory,
		storeRepositoryFactory: storeRepositoryFactory,
	}
}

func (sc *StoreController) GetStores(c echo.Context) error {
	return sc.newStoreInputPort(c).GetStores()
}

func (sc *StoreController) GetNearStores(c echo.Context) error {
	return sc.newStoreInputPort(c).GetNearStores()
}

func (sc *StoreController) SaveFavoriteStore(c echo.Context) error {
	return sc.newStoreInputPort(c).SaveFavoriteStore()
}

/* ここでpresenterにecho.Contextを渡している！起爆！！！（遅延） */
/* これによって、presenterのinterface(outputport)にecho.Contextを書かなくて良くなる */
func (sc *StoreController) newStoreInputPort(c echo.Context) port.StoreInputPort {
	storeOutputPort := sc.storeOutputFactory(c)
	storeDriver := sc.storeDriverFactory
	googleMapDriver := sc.googleMapDriverFactory
	storeRepository := sc.storeRepositoryFactory(storeDriver, googleMapDriver)
	return sc.storeInputFactory(storeRepository, storeOutputPort)
}
