package presenter

import (
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StorePresenter struct {
	c echo.Context
}

func NewStoreOutputPort(c echo.Context) port.StoreOutputPort {
	return &StorePresenter{c: c}
}

type StoreOutputJson struct {
	Stores  []storeForPresenter `json:"stores"`
}

type storeForPresenter struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (sp *StorePresenter) OutputAllStores(stores []*model.Store) error {
	json_stores := make([]storeForPresenter, 0)
	for _, v := range stores {
		json_stores = append(json_stores, storeForPresenter{
			Id:   v.Id,
			Name: v.Name,
		})
	}
	output_json := &StoreOutputJson{Stores: json_stores}
	return sp.c.JSON(http.StatusOK, output_json)
}
