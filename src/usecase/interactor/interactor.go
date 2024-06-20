package interactor

import (
	port "clean-storemap-api/src/usecase/port"
)

type StoreInteractor struct {
	storeRepository port.StoreRepository
	storeOutputPort port.StoreOutputPort
}

// type StoreInteractor struct{}

// 本来は引数としてOutputPortが必要
func NewStoreInputPort(storeRepository port.StoreRepository, storeOutputPort port.StoreOutputPort) port.StoreInputPort {
	return &StoreInteractor{
		storeRepository: storeRepository,
		storeOutputPort: storeOutputPort,
	}
}

// func NewStoreInputPort() port.StoreInputPort {
// 	return &StoreInteractor{}
// }

func (si *StoreInteractor) GetStores() error {
	stores, err := si.storeRepository.GetAll()
	if err != nil {
		return err
	}
	return si.storeOutputPort.OutputAllStores(stores)
	// dummyStores := make([]*model.Store, 0)
	// dummyStores = append(dummyStores, &model.Store{Id: 1, Name: "aa"})
	// stores := dummyStores
	// return stores, nil
}
