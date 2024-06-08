package controller

import (
	model "clean-storemap-api/src/entity"
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

type StoreInputPort interface {
	GetStores() ([]model.Store, error)
}

type StoreController struct {
	StoreInputPort StoreInputPort
}

func (sc *StoreController) GetStores(c echo.Context) error {
	stores, err := sc.StoreInputPort.GetStores()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
	}
	json_stores := make([]storeForController, 0)
	for _, v := range stores {
		json_stores = append(json_stores, storeForController{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	stores_json := &StoreJson{ResponseCode: 200, Message: "iine", Stores: json_stores}
	return c.JSON(http.StatusOK, stores_json)
}
