package gateway

import (
	db "clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
)

// Dbの要素を構造体として渡す必要がある。
type StoreGateway struct {
	storeDriver StoreDriver
}
type StoreDriver interface {
	GetStores() ([]*db.Store, error)
}

func NewStoreRepository(storeDriver StoreDriver) port.StoreRepository {
	return &StoreGateway{
		storeDriver: storeDriver,
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
			Id:   v.Id,
			Name: v.Name,
		})
	}
	return stores, nil
}
