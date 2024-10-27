package port

import (
	model "clean-storemap-api/src/entity"
)

type StoreInputPort interface {
	GetStores() error
	GetStoresOpeningHours() error
}

type StoreRepository interface {
	GetAll() ([]*model.Store, error)
	GetOpeningHours() ([]*model.Store, error)
}

type StoreOutputPort interface {
	OutputAllStores([]*model.Store) error
}
