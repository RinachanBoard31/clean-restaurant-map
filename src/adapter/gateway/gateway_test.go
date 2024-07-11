package gateway

import (
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func makeDummyStores() ([]*db.Store, error) {
	dummyStores := make([]*db.Store, 0)
	dummyStores = append(dummyStores, &db.Store{Id: 1, Name: "サイゼリヤ"})
	dummyStores = append(dummyStores, &db.Store{Id: 2, Name: "ガスト"})
	return dummyStores, nil
}

type MockStoreRepository struct {
	mock.Mock
}

func (m *MockStoreRepository) GetStores() ([]*db.Store, error) {
	args := m.Called()
	return args.Get(0).([]*db.Store), args.Error(1)
}

func TestGetAll(t *testing.T) {
	/* Arrange */
	sg := &StoreGateway{}
	mockStoreRepository := new(MockStoreRepository)
	mockStoreRepository.On("GetStores").Return(makeDummyStores())
	sg.storeDriver = mockStoreRepository
	stores := make([]*model.Store, 0)
	stores = append(stores, &model.Store{Id: 1, Name: "サイゼリヤ"})
	stores = append(stores, &model.Store{Id: 2, Name: "ガスト"})
	expected := stores

	/* Act */
	actual, _ := sg.GetAll()

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expected, actual)
	// storeDriver.GetStores()が1回呼ばれること
	mockStoreRepository.AssertNumberOfCalls(t, "GetStores", 1)
}
