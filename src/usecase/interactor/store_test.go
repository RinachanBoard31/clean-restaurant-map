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

func (m *MockStoreRepository) GetNearStores() ([]*model.Store, error) {
	args := m.Called()
	return args.Get(0).([]*model.Store), args.Error(1)
}

func (m *MockStoreRepository) SaveFavoriteStore(store *model.Store, userId string) error {
	args := m.Called(store, userId)
	return args.Error(0)
}

func (m *MockStoreOutputPort) OutputAllStores(stores []*model.Store) error {
	args := m.Called(stores)
	return args.Error(0)
}

func (m *MockStoreOutputPort) OutputSaveFavoriteStoreResult() error {
	args := m.Called()
	return args.Error(0)
}

func TestGetStores(t *testing.T) {
	/* Arrange */
	expected := errors.New("")
	stores := make([]*model.Store, 0)
	stores = append(
		stores,
		&model.Store{
			Id:                  "Id001",
			Name:                "UEC cafe",
			RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
			PriceLevel:          "PRICE_LEVEL_MODERATE",
			Location:            model.Location{Lat: "35.713", Lng: "139.762"},
		},
	)

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

func TestGetNearStores(t *testing.T) {
	/* Arrange */
	expected := errors.New("")
	stores := make([]*model.Store, 0)
	stores = append(
		stores,
		&model.Store{
			Id:                  "Id001",
			Name:                "UEC cafe",
			RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
			PriceLevel:          "PRICE_LEVEL_MODERATE",
			Location:            model.Location{Lat: "35.713", Lng: "139.762"},
		},
	)

	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("GetNearStores").Return(stores, nil)
	mockStoreOutputPort := new(MockStoreOutputPort)
	mockStoreOutputPort.On("OutputAllStores", stores).Return(expected)

	si := &StoreInteractor{storeRepository: mockStoreRepository, storeOutputPort: mockStoreOutputPort}

	/* Act */
	actual := si.GetNearStores()

	/* Assert */
	assert.Equal(t, expected, actual)
	mockStoreRepository.AssertNumberOfCalls(t, "GetNearStores", 1)
	mockStoreOutputPort.AssertNumberOfCalls(t, "OutputAllStores", 1)
	mockStoreOutputPort.AssertCalled(t, "OutputAllStores", stores)
}

func TestSaveFavoriteStore(t *testing.T) {
	/* Arrange */
	var expected error = nil
	store := &model.Store{
		Id:                  "Id001",
		Name:                "UEC cafe",
		RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
		PriceLevel:          "PRICE_LEVEL_MODERATE",
		Location:            model.Location{Lat: "35.713", Lng: "139.762"},
	}
	userId := "test@example.com"

	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("SaveFavoriteStore", store, userId).Return(nil)
	mockStoreOutputPort := new(MockStoreOutputPort)
	mockStoreOutputPort.On("OutputSaveFavoriteStoreResult").Return(nil)

	si := &StoreInteractor{storeRepository: mockStoreRepository, storeOutputPort: mockStoreOutputPort}

	/* Act */
	actual := si.SaveFavoriteStore(store, userId)

	/* Assert */
	assert.Equal(t, expected, actual)
	mockStoreRepository.AssertNumberOfCalls(t, "SaveFavoriteStore", 1)
	mockStoreOutputPort.AssertNumberOfCalls(t, "OutputSaveFavoriteStoreResult", 1)
	mockStoreRepository.AssertCalled(t, "SaveFavoriteStore", store, userId)
}
