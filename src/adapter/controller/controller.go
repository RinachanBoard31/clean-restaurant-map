package controller

import (
	model "clean-storemap-api/src/entity"
	//"clean-storemap-api/src/usecase/port"
	// "errors"
	"net/http"

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

//	type StoreController struct {
//		storeInputPort port.StoreInputPort
//	}
type StoreController struct{}

// func NewStoreController(storeInputPort port.StoreInputPort) StoreI {
// 	return &StoreController{
// 		storeInputPort: storeInputPort,
// 	}
// }

func NewStoreController() StoreI {
	return &StoreController{}
}

func (sc *StoreController) GetStores(c echo.Context) error {
	//stores, err := sc.storeInputPort.GetStores() <-本当はこれを呼ぶ
	dummyStores := make([]model.Store, 0)
	dummyStores = append(dummyStores, model.Store{Id: 1, Name: "aa"})

	// err := errors.New("")
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	// }
	json_stores := make([]storeForController, 0)
	for _, v := range dummyStores {
		json_stores = append(json_stores, storeForController{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	stores_json := &StoreJson{ResponseCode: 200, Message: "iine", Stores: json_stores}
	return c.JSON(http.StatusOK, stores_json)
}
