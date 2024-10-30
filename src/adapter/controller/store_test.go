package controller

import (
	"clean-storemap-api/src/adapter/gateway"
	api "clean-storemap-api/src/driver/api"
	db "clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStoreDriverFactory struct {
	mock.Mock
}

type MockGoogleMapDriverFactory struct {
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

func (m *MockGoogleMapDriverFactory) GetStores() ([]*api.Store, error) {
	args := m.Called()
	return args.Get(0).([]*api.Store), args.Error(1)
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

func (m *MockStoreRepositoryFactoryFuncObject) GetNearStores() ([]*model.Store, error) {
	args := m.Called()
	return args.Get(0).([]*model.Store), args.Error(1)
}

func mockStoreRepositoryFactoryFunc(storeDriver gateway.StoreDriver, googleMapDriver gateway.GoogleMapDriver) port.StoreRepository {
	return &MockStoreRepositoryFactoryFuncObject{}
}

func (m *MockStoreInputFactoryFuncObject) GetStores() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockStoreInputFactoryFuncObject) GetNearStores() error {
	args := m.Called()
	return args.Error(0)
}

// Validationのために必要なメソッド
type CustomValidator struct {
	validator *validator.Validate
}

func NewValidator() echo.Validator {
	return &CustomValidator{validator: validator.New()}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func newRouter() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	e.Validator = NewValidator()
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

func TestGetNearStores(t *testing.T) {
	/* Arrange */
	c, rec := newRouter()
	expected := errors.New("")

	mockGoogleMapDriverFactory := new(MockGoogleMapDriverFactory)
	mockGoogleMapDriverFactory.On("GetStores").Return(expected)

	sc := &StoreController{
		googleMapDriverFactory: mockGoogleMapDriverFactory,
		storeOutputFactory:     mockStoreOutputFactoryFunc,
		storeRepositoryFactory: mockStoreRepositoryFactoryFunc,
	}

	mockStoreInputFactoryFuncObject := new(MockStoreInputFactoryFuncObject)
	mockStoreInputFactoryFuncObject.On("GetNearStores").Return(expected)
	sc.storeInputFactory = func(repository port.StoreRepository, output port.StoreOutputPort) port.StoreInputPort {
		return mockStoreInputFactoryFuncObject
	}

	/* Act */
	actual := sc.GetNearStores(c)

	/* Assert */
	assert.Equal(t, expected, actual)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockStoreInputFactoryFuncObject.AssertNumberOfCalls(t, "GetNearStores", 1)
}
