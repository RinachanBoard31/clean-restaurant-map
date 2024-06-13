package interactor

import (
	model "clean-storemap-api/src/entity"
	port "clean-storemap-api/src/usecase/port"
)

//	type StoreInteractor struct {
//		storeRepository port.StoreRepository
//	}
type StoreInteractor struct{}

// 本来は引数としてOutputPortが必要
// func NewStoreInputPort(storeRepository port.StoreRepository) port.StoreInputPort {
// 	return &StoreInteractor{
// 		storeRepository: storeRepository,
// 	}
// }

func NewStoreInputPort() port.StoreInputPort {
	return &StoreInteractor{}
}

func (si *StoreInteractor) GetStores() ([]*model.Store, error) {
	// stores, err := si.storeRepository.GetAll()  <-本当はこれを呼ぶ
	dummyStores := make([]*model.Store, 0)
	dummyStores = append(dummyStores, &model.Store{Id: 1, Name: "aa"})
	stores := dummyStores

	return stores, nil
}
