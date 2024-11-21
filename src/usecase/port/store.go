package port

import (
	model "clean-storemap-api/src/entity"
)

type StoreInputPort interface {
	GetStores() error
	GetNearStores() error
	SaveFavoriteStore(store *model.Store, userId string) error
}

type StoreRepository interface {
	GetAll() ([]*model.Store, error)
	GetNearStores() ([]*model.Store, error)
	SaveFavoriteStore(store *model.Store, userId string) error
}

type StoreOutputPort interface {
	OutputAllStores([]*model.Store) error
	OutputSaveFavoriteStoreResult() error
}
