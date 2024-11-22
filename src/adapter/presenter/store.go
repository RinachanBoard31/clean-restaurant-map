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
	Stores []storeForPresenter `json:"stores"`
}

type locationForPresenter struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type storeForPresenter struct {
	Id                  string               `json:"id"`
	Name                string               `json:"name"`
	RegularOpeningHours string               `json:"regularOpeningHours"`
	PriceLevel          string               `json:"priceLevel"`
	Location            locationForPresenter `json:"location"`
}

func (sp *StorePresenter) OutputAllStores(stores []*model.Store) error {
	json_stores := make([]storeForPresenter, 0)
	for _, v := range stores {
		json_stores = append(json_stores, storeForPresenter{
			Id:                  v.Id,
			Name:                v.Name,
			RegularOpeningHours: v.RegularOpeningHours,
			PriceLevel:          v.PriceLevel,
			Location: locationForPresenter{
				Latitude:  v.Location.Lat,
				Longitude: v.Location.Lng,
			},
		})
	}
	output_json := &StoreOutputJson{Stores: json_stores}
	return sp.c.JSON(http.StatusOK, output_json)
}

func (sp *StorePresenter) OutputSaveFavoriteStoreResult() error {
	return sp.c.JSON(http.StatusOK, map[string]interface{}{})
}

func (sp *StorePresenter) OutputAlreadyExistFavorite() error {
	errMsg := "Already exist favorite store"
	return sp.c.JSON(http.StatusConflict, map[string]interface{}{"error": errMsg})
}
