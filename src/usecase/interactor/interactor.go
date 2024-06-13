package interactor

import (
	model "clean-storemap-api/src/entity"
	port "clean-storemap-api/src/usecase/port"
)

type StoreInteractor struct {
	storeRepository port.StoreRepository
}

func (si *StoreInteractor) GetStores() ([]*model.Store, error) {
	stores, err := si.storeRepository.GetAll()

	return stores, err
}
