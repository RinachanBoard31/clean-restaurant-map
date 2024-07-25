package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/adapter/gateway"
	"clean-storemap-api/src/adapter/presenter"
	"clean-storemap-api/src/driver/db"
	"clean-storemap-api/src/usecase/interactor"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"
)

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

func ActivateRouter() {
	e := echo.New()
	e.Validator = NewValidator()
	db.InitDB()

	storeDriver := NewDriverFactory()
	storeOutputPort := NewOutputFactory()
	storeInputPort := NewInputFactory()
	storeRepository := NewRepositoryFactory()
	store_controller := controller.NewStoreController(storeDriver, storeOutputPort, storeInputPort, storeRepository)

	e.GET("/", store_controller.GetStores)

	e.Logger.Fatal(e.Start(":8080"))
}

func NewDriverFactory() controller.DriverFactory {
	return &db.DbStoreDriver{}
}

func NewOutputFactory() controller.OutputFactory {
	return presenter.NewStoreOutputPort
}

func NewInputFactory() controller.InputFactory {
	return interactor.NewStoreInputPort
}

func NewRepositoryFactory() controller.RepositoryFactory {
	return gateway.NewStoreRepository
}
