package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/adapter/gateway"
	"clean-storemap-api/src/adapter/presenter"
	"clean-storemap-api/src/usecase/interactor"
	"net/http"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

func newRouter() (*echo.Echo, echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return e, c, rec
}

func ActivateRouter() {
	e, c, _ := newRouter()

	// コンストラクタを呼び出していないのでエラーになる。
	// RepositoryとOutputの実装が完了しないとエラーは解消しない。
	storeRepository := gateway.NewStoreRepository()
	storeOutputPort := presenter.NewStoreOutputPort(c)
	inputPort := interactor.NewStoreInputPort(storeRepository, storeOutputPort)
	controller := controller.NewStoreController(inputPort)

	e.GET("/", controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}
