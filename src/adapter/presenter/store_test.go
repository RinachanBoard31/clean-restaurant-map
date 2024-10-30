package presenter

import (
	model "clean-storemap-api/src/entity"
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"

	"github.com/labstack/echo/v4"
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
	expected := "{\"stores\":[{\"id\":\"Id001\",\"name\":\"UEC cafe\",\"regularOpeningHours\":\"Sat: 06:00 - 22:00, Sun: 06:00 - 22:00\",\"priceLevel\":\"PRICE_LEVEL_MODERATE\"}]}\n"
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
