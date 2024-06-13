package controller

import (
	model "clean-storemap-api/src/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStoreInputPort struct {
	mock.Mock
}

func (m *MockStoreInputPort) GetStores() ([]*model.Store, error) {
	args := m.Called()
	return args.Get(0).([]*model.Store), args.Error(1)
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
	expected := "{\"ResponseCode\":200,\"Message\":\"iine\",\"Stores\":[{\"Id\":1,\"Name\":\"aa\"}]}\n"
	stores := make([]*model.Store, 0)
	stores = append(stores, &model.Store{Id: 1, Name: "aa"})

	// InputPortのGetStoresのモックを作成
	mockStoreInputPort := new(MockStoreInputPort)
	mockStoreInputPort.On("GetStores").Return(stores, nil)
	sc := &StoreController{storeInputPort: mockStoreInputPort}

	/* Act */
	actual := sc.GetStores(c)

	/* Assert */
	// 指定したStatusCodeが返却されること
	if assert.NoError(t, actual) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
	// 指定したResponseBodyが返却されること
	if assert.NoError(t, actual) {
		assert.Equal(t, expected, rec.Body.String())
	}
	// InputPortのGetStoresが1回呼ばれること
	mockStoreInputPort.AssertNumberOfCalls(t, "GetStores", 1)
}
