package gateway

import (
	api "clean-storemap-api/src/driver/api"
	db "clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeDummyDbStores() ([]*db.Store, error) {
	dummyStores := make([]*db.Store, 0)
	dummyStores = append(dummyStores, &db.Store{
		Id:                  "Id001",
		Name:                "UEC cafe",
		RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
		PriceLevel:          "PRICE_LEVEL_MODERATE",
	})
	dummyStores = append(dummyStores, &db.Store{
		Id:                  "Id002",
		Name:                "UEC restaurant",
		RegularOpeningHours: "Sat: 11:00 - 20:00, Sun: 11:00 - 20:00",
		PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
	})
	return dummyStores, nil
}

func makeDummyApiStores() ([]*api.Store, error) {
	dummyStores := make([]*api.Store, 0)
	dummyStores = append(dummyStores, &api.Store{
		Id:                  "Id001",
		Name:                "UEC cafe",
		RegularOpeningHours: []string{"Sat: 06:00 - 22:00", "Sun: 06:00 - 22:00"},
		PriceLevel:          "PRICE_LEVEL_MODERATE",
	})
	dummyStores = append(dummyStores, &api.Store{
		Id:                  "Id002",
		Name:                "UEC restaurant",
		RegularOpeningHours: []string{"Sat: 11:00 - 20:00", "Sun: 11:00 - 20:00"},
		PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
	})
	return dummyStores, nil
}

type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) GetStores() ([]*db.Store, error) {
	args := m.Called()
	return args.Get(0).([]*db.Store), args.Error(1)
}

type MockGoogleMapRepository struct {
	mock.Mock
}

func (m *MockGoogleMapRepository) GetStores() ([]*api.Store, error) {
	args := m.Called()
	return args.Get(0).([]*api.Store), args.Error(1)
}

func TestGetAll(t *testing.T) {
	/* Arrange */
	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("GetStores").Return(makeDummyDbStores())
	sg := &StoreGateway{storeDriver: mockStoreRepository}
	stores := make([]*model.Store, 0)
	stores = append(
		stores,
		&model.Store{
			Id:                  "Id001",
			Name:                "UEC cafe",
			RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
			PriceLevel:          "PRICE_LEVEL_MODERATE",
		},
	)
	stores = append(
		stores,
		&model.Store{
			Id:                  "Id002",
			Name:                "UEC restaurant",
			RegularOpeningHours: "Sat: 11:00 - 20:00, Sun: 11:00 - 20:00",
			PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
		},
	)
	expected := stores

	/* Act */
	actual, _ := sg.GetAll()

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expected, actual)
	// storeDriver.GetStores()が1回呼ばれること
	mockStoreRepository.AssertNumberOfCalls(t, "GetStores", 1)
}

func TestGetOpeningHours(t *testing.T) {
	/* Arrange */
	mockGoogleMapRepository := new(MockGoogleMapRepository)
	mockGoogleMapRepository.On("GetStores").Return(makeDummyApiStores())
	sg := &StoreGateway{googleMapDriver: mockGoogleMapRepository}
	stores := make([]*model.Store, 0)
	stores = append(
		stores,
		&model.Store{
			Id:                  "Id001",
			Name:                "UEC cafe",
			RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
			PriceLevel:          "PRICE_LEVEL_MODERATE",
		},
		&model.Store{
			Id:                  "Id002",
			Name:                "UEC restaurant",
			RegularOpeningHours: "Sat: 11:00 - 20:00, Sun: 11:00 - 20:00",
			PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
		},
	)
	expected := stores

	/* Act */
	actual, _ := sg.GetOpeningHours()

	/* Assert */
	assert.Equal(t, expected, actual)
	mockGoogleMapRepository.AssertNumberOfCalls(t, "GetStores", 1)
}
