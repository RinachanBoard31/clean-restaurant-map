package interactor

import (
	model "clean-storemap-api/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) GetAll() ([]*model.Store, error) {
	args := m.Called()
	return args.Get(0).([]*model.Store), args.Error(1)
}

func TestGetStores(t *testing.T) {
	expected := make([]*model.Store, 0)
	expected = append(expected, &model.Store{Id: 1, Name: "interactor"})
	stores := make([]*model.Store, 0)
	stores = append(stores, &model.Store{Id: 1, Name: "interactor"})

	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("GetAll").Return(stores, nil)
	si := &StoreInteractor{storeRepository: mockStoreRepository}

	actual, _ := si.GetStores()

	/* Assert */
	// Repositoryから返ってくる構造体がドメインモデルであること
	assert.Equal(t, expected, actual)

	// RepositoryのGetAllが1回呼ばれること
	mockStoreRepository.AssertNumberOfCalls(t, "GetAll", 1)
}
