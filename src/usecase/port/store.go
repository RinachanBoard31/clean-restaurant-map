package port

import (
	model "clean-storemap-api/src/entity"
)

type StoreInputPort interface {
	GetStores() error
	GetNearStores() error
	SaveFavoriteStore(store *model.Store, userId int) error
	GetTopFavoriteStores() error
}

type StoreRepository interface {
	GetAll() ([]*model.Store, error)
	GetNearStores() ([]*model.Store, error)
	ExistFavorite(store *model.Store, userId int) (bool, error)
	SaveFavoriteStore(store *model.Store, userId int) error
	GetTopFavoriteStores() ([]*model.Store, error)
}

type StoreOutputPort interface {
	OutputAllStores([]*model.Store) error
	OutputSaveFavoriteStoreResult() error
	OutputAlreadyExistFavorite() error
}
