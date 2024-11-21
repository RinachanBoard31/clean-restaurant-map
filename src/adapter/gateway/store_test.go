package gateway

import (
	api "clean-storemap-api/src/driver/api"
	db "clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeDummyDbStores() ([]*db.FavoriteStore, error) {
	dummyStores := make([]*db.FavoriteStore, 0)
	dummyStores = append(dummyStores, &db.FavoriteStore{
		Id:                  "Id001",
		Name:                "UEC cafe",
		RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
		PriceLevel:          "PRICE_LEVEL_MODERATE",
		Latitude:            "35.713",
		Longitude:           "139.762",
	})
	dummyStores = append(dummyStores, &db.FavoriteStore{
		Id:                  "Id002",
		Name:                "UEC restaurant",
		RegularOpeningHours: "Sat: 11:00 - 20:00, Sun: 11:00 - 20:00",
		PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
		Latitude:            "35.714",
		Longitude:           "139.763",
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
		Location:            api.Location{Lat: 35.713, Lng: 139.762},
	})
	dummyStores = append(dummyStores, &api.Store{
		Id:                  "Id002",
		Name:                "UEC restaurant",
		RegularOpeningHours: []string{"Sat: 11:00 - 20:00", "Sun: 11:00 - 20:00"},
		PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
		Location:            api.Location{Lat: 35.714, Lng: 139.763},
	})
	return dummyStores, nil
}

type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) GetStores() ([]*db.FavoriteStore, error) {
	args := m.Called()
	return args.Get(0).([]*db.FavoriteStore), args.Error(1)
}

func (m *MockStoreRepository) SaveStore(*db.FavoriteStore) error {
	args := m.Called()
	return args.Error(0)
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
			Location:            model.Location{Lat: "35.713", Lng: "139.762"},
		},
	)
	stores = append(
		stores,
		&model.Store{
			Id:                  "Id002",
			Name:                "UEC restaurant",
			RegularOpeningHours: "Sat: 11:00 - 20:00, Sun: 11:00 - 20:00",
			PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
			Location:            model.Location{Lat: "35.714", Lng: "139.763"},
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

func TestGetNearStores(t *testing.T) {
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
			Location:            model.Location{Lat: "35.713000", Lng: "139.762000"},
		},
		&model.Store{
			Id:                  "Id002",
			Name:                "UEC restaurant",
			RegularOpeningHours: "Sat: 11:00 - 20:00, Sun: 11:00 - 20:00",
			PriceLevel:          "PRICE_LEVEL_INEXPENSIVE",
			Location:            model.Location{Lat: "35.714000", Lng: "139.763000"},
		},
	)
	expected := stores

	/* Act */
	actual, _ := sg.GetNearStores()

	/* Assert */
	assert.Equal(t, expected, actual)
	mockGoogleMapRepository.AssertNumberOfCalls(t, "GetStores", 1)
}

func TestSaveFavoriteStore(t *testing.T) {
	/* Arrange */
	var expected error = nil
	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("SaveStore").Return(nil)
	sg := &StoreGateway{storeDriver: mockStoreRepository}
	store := &model.Store{
		Id:                  "Id001",
		Name:                "UEC cafe",
		RegularOpeningHours: "Sat: 06:00 - 22:00, Sun: 06:00 - 22:00",
		PriceLevel:          "PRICE_LEVEL_MODERATE",
		Location:            model.Location{Lat: "35.713", Lng: "139.762"},
	}

	/* Act */
	actual := sg.SaveFavoriteStore(store)

	/* Assert */
	assert.Equal(t, expected, actual)
	mockStoreRepository.AssertNumberOfCalls(t, "SaveStore", 1)
}
