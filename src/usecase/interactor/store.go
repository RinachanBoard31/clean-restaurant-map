package interactor

import (
	model "clean-storemap-api/src/entity"
	port "clean-storemap-api/src/usecase/port"
)

type StoreInteractor struct {
	storeRepository port.StoreRepository
	storeOutputPort port.StoreOutputPort
}

func NewStoreInputPort(storeRepository port.StoreRepository, storeOutputPort port.StoreOutputPort) port.StoreInputPort {
	return &StoreInteractor{
		storeRepository: storeRepository,
		storeOutputPort: storeOutputPort,
	}
}

func (si *StoreInteractor) GetStores() error {
	stores, err := si.storeRepository.GetAll()
	if err != nil {
		return err
	}
	return si.storeOutputPort.OutputAllStores(stores)
}

func (si *StoreInteractor) GetNearStores() error {
	places, err := si.storeRepository.GetNearStores()
	if err != nil {
		return err
	}
	return si.storeOutputPort.OutputAllStores(places)
}

func (si *StoreInteractor) SaveFavoriteStore(store *model.Store, userId string) error {
	if err := si.storeRepository.SaveFavoriteStore(store, userId); err != nil {
		return err
	}
	if err := si.storeOutputPort.OutputSaveFavoriteStoreResult(); err != nil {
		return err
	}
	return nil
}
