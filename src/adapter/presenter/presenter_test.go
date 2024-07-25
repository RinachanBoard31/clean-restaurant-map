package presenter

import (
	model "clean-storemap-api/src/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func newRouter() (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestOutputAllStores(t *testing.T) {
	/* Arrange */
	expected := "{\"stores\":[{\"id\":1,\"name\":\"aa\"}]}\n"
	stores := make([]*model.Store, 0)
	stores = append(stores, &model.Store{Id: 1, Name: "aa"})
	c, rec := newRouter()
	sp := &StorePresenter{c: c}

	/* Act */
	actual := sp.OutputAllStores(stores)

	/* Assert */
	// sp.OutputAllStores()がJSONを返すこと
	if assert.NoError(t, actual) {
		assert.Equal(t, expected, rec.Body.String())
	}
}
