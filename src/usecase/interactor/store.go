package interactor

import (
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

func (si *StoreInteractor) GetStoresOpeningHours() error {
	places, err := si.storeRepository.GetOpeningHours()
	if err != nil {
		return err
	}
	return si.storeOutputPort.OutputAllStores(places)
}
