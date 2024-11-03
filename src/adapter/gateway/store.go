package gateway

import (
	api "clean-storemap-api/src/driver/api"
	db "clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
	"fmt"
	"strings"
)

// Dbの要素を構造体として渡す必要がある。
type StoreGateway struct {
	storeDriver     StoreDriver
	googleMapDriver GoogleMapDriver
}

type StoreDriver interface {
	GetStores() ([]*db.Store, error)
}

type GoogleMapDriver interface {
	GetStores() ([]*api.Store, error)
}

func NewStoreRepository(storeDriver StoreDriver, googleMapDriver GoogleMapDriver) port.StoreRepository {
	return &StoreGateway{
		storeDriver:     storeDriver,
		googleMapDriver: googleMapDriver,
	}
}

func (sg *StoreGateway) GetAll() ([]*model.Store, error) {
	dbStores, err := sg.storeDriver.GetStores()
	if err != nil {
		return nil, err
	}
	stores := make([]*model.Store, 0)
	for _, v := range dbStores {
		stores = append(stores, &model.Store{
			Id:                  v.Id,
			Name:                v.Name,
			RegularOpeningHours: v.RegularOpeningHours,
			PriceLevel:          v.PriceLevel,
			Location: model.Location{
				Lat: v.Latitude,
				Lng: v.Longitude,
			},
		})
	}
	return stores, nil
}

func (sg *StoreGateway) GetNearStores() ([]*model.Store, error) {
	apiStores, err := sg.googleMapDriver.GetStores()
	if err != nil {
		return nil, err
	}
	stores := make([]*model.Store, 0)
	for _, v := range apiStores {
		stores = append(stores, &model.Store{
			Id:                  v.Id,
			Name:                v.Name,
			RegularOpeningHours: strings.Join(v.RegularOpeningHours, ", "),
			PriceLevel:          v.PriceLevel,
			Location: model.Location{
				Lat: fmt.Sprintf("%f", v.Location.Lat),
				Lng: fmt.Sprintf("%f", v.Location.Lng),
			},
		})
	}
	return stores, nil
}
