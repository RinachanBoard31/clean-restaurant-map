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

	// StoreControllerのDI
	storeDriver := NewStoreDriverFactory()
	storeOutputPort := NewStoreOutputFactory()
	storeInputPort := NewStoreInputFactory()
	storeRepository := NewStoreRepositoryFactory()
	storeController := controller.NewStoreController(storeDriver, storeOutputPort, storeInputPort, storeRepository)

	// UserControllerのDI
	userDriver := NewUserDriverFactory()
	userOutputPort := NewUserOutputFactory()
	userInputPort := NewUserInputFactory()
	userRepository := NewUserRepositoryFactory()
	userController := controller.NewUserController(userDriver, userOutputPort, userInputPort, userRepository)

	e.GET("/", storeController.GetStores)
	e.POST("/user", userController.CreateUser)

	e.Logger.Fatal(e.Start(":8080"))
}

// StoreのDI
func NewStoreDriverFactory() controller.StoreDriverFactory {
	return &db.DbStoreDriver{}
}

func NewStoreOutputFactory() controller.StoreOutputFactory {
	return presenter.NewStoreOutputPort
}

func NewStoreInputFactory() controller.StoreInputFactory {
	return interactor.NewStoreInputPort
}

func NewStoreRepositoryFactory() controller.StoreRepositoryFactory {
	return gateway.NewStoreRepository
}

// UserのDI
func NewUserDriverFactory() controller.UserDriverFactory {
	return &db.DbUserDriver{}
}

func NewUserOutputFactory() controller.UserOutputFactory {
	return presenter.NewUserOutputPort
}

func NewUserInputFactory() controller.UserInputFactory {
	return interactor.NewUserInputPort
}

func NewUserRepositoryFactory() controller.UserRepositoryFactory {
	return gateway.NewUserRepository
}
