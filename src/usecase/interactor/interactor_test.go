package interactor

import (
	model "clean-storemap-api/src/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStoreRepository struct {
	mock.Mock
}
type MockStoreOutputPort struct {
	mock.Mock
}

func (m *MockStoreRepository) GetAll() ([]*model.Store, error) {
	args := m.Called()
	return args.Get(0).([]*model.Store), args.Error(1)
}

func (m *MockStoreOutputPort) OutputAllStores(stores []*model.Store) error {
	args := m.Called(stores)
	return args.Error(0)
}

func TestGetStores(t *testing.T) {
	/* Arrange */
	/* ### Repository テスト用 #################
	expected := make([]*model.Store, 0)
	expected = append(expected, &model.Store{Id: 1, Name: "interactor"})
	########################################## */
	expected := errors.New("")
	stores := make([]*model.Store, 0)
	stores = append(stores, &model.Store{Id: 1, Name: "interactor"})

	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("GetAll").Return(stores, nil)
	mockStoreOutputPort := new(MockStoreOutputPort)
	mockStoreOutputPort.On("OutputAllStores", stores).Return(expected)

	si := &StoreInteractor{storeRepository: mockStoreRepository, storeOutputPort: mockStoreOutputPort}

	/* Act */
	actual := si.GetStores()

	/* Assert */
	// GetStores()がOutputAllStores()を返すこと
	assert.Equal(t, expected, actual)
	// RepositoryのGetAllが1回呼ばれること
	mockStoreRepository.AssertNumberOfCalls(t, "GetAll", 1)
	// OutputPortのOutputAllStoresが1回呼ばれること
	mockStoreOutputPort.AssertNumberOfCalls(t, "OutputAllStores", 1)
	// OutputPortのOutputAllStoresがstoresを引数として呼ばれること
	mockStoreOutputPort.AssertCalled(t, "OutputAllStores", stores)
}
