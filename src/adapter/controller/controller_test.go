package controller

import (
	"clean-storemap-api/src/adapter/gateway"
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStoreDriverFactory struct {
	mock.Mock
}

type MockStoreOutputFactoryFuncObject struct {
	mock.Mock
}

type MockStoreRepositoryFactoryFuncObject struct {
	mock.Mock
}

type MockStoreInputFactoryFuncObject struct {
	mock.Mock
}

func (m *MockStoreDriverFactory) GetStores() ([]*db.Store, error) {
	args := m.Called()
	return args.Get(0).([]*db.Store), args.Error(1)
}

func (m *MockStoreOutputFactoryFuncObject) OutputAllStores([]*model.Store) error {
	args := m.Called()
	return args.Error(0)
}

func mockStoreOutputFactoryFunc(c echo.Context) port.StoreOutputPort {
	return &MockStoreOutputFactoryFuncObject{}
}

func (m *MockStoreRepositoryFactoryFuncObject) GetAll() ([]*model.Store, error) {
	args := m.Called()
	return args.Get(0).([]*model.Store), args.Error(1)
}

func mockStoreRepositoryFactoryFunc(storeDriver gateway.StoreDriver) port.StoreRepository {
	return &MockStoreRepositoryFactoryFuncObject{}
}

func (m *MockStoreInputFactoryFuncObject) GetStores() error {
	args := m.Called()
	return args.Error(0)
}

func newRouter() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}

func TestGetStores(t *testing.T) {
	/* Arrange */
	c, rec := newRouter()
	expected := errors.New("")

	// Driverだけは実体が必要
	mockStoreDriverFactory := new(MockStoreDriverFactory)
	mockStoreDriverFactory.On("GetStores").Return([]*db.Store{}, nil)

	// InputPortのGetStoresのモックを作成
	sc := &StoreController{
		storeDriverFactory:     mockStoreDriverFactory,
		storeOutputFactory:     mockStoreOutputFactoryFunc,
		storeRepositoryFactory: mockStoreRepositoryFactoryFunc,
	}

	// newStoreInputPort.GetStores()をするためには、GetStores()を持つmockStoreInputFactoryFuncObjectがstoreInputFactoryに必要だから無名関数でreturnする必要があった
	mockStoreInputFactoryFuncObject := new(MockStoreInputFactoryFuncObject)
	mockStoreInputFactoryFuncObject.On("GetStores").Return(expected)
	sc.storeInputFactory = func(repository port.StoreRepository, output port.StoreOutputPort) port.StoreInputPort {
		return mockStoreInputFactoryFuncObject
	}

	/* Act */
	actual := sc.GetStores(c)

	/* Assert */
	// sc.GetStores()がstoreInputPort.GetStores()を返すこと
	assert.Equal(t, expected, actual)
	// echoが正しく起動したか
	assert.Equal(t, http.StatusOK, rec.Code)

	// InputPortのGetStoresが1回呼ばれること
	mockStoreInputFactoryFuncObject.AssertNumberOfCalls(t, "GetStores", 1)
}
