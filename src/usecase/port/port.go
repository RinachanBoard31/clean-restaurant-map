package port

import model "clean-storemap-api/src/entity"


type StoreInputPort interface {
	GetStores() ([]*model.Store, error)
}
