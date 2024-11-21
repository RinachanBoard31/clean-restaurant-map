package port

import (
	model "clean-storemap-api/src/entity"
)

type StoreInputPort interface {
	GetStores() error
	GetNearStores() error
	SaveFavoriteStore(*model.Store) error
}

type StoreRepository interface {
	GetAll() ([]*model.Store, error)
	GetNearStores() ([]*model.Store, error)
	SaveFavoriteStore(*model.Store) error
}

type StoreOutputPort interface {
	OutputAllStores([]*model.Store) error
	OutputSaveFavoriteStoreResult() error
}
