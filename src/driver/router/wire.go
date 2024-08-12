//go:build wireinject
// +build wireinject

package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/adapter/gateway"
	"clean-storemap-api/src/adapter/presenter"
	"clean-storemap-api/src/driver/db"
	"clean-storemap-api/src/usecase/interactor"
	"context"

	"github.com/google/wire"
	"github.com/labstack/echo/v4"
)

var echoSet = wire.NewSet(
	NewEcho,
)

var driverSet = wire.NewSet(
	NewStoreDriverFactory,
	NewUserDriverFactory,
)

var inputPortSet = wire.NewSet(
	NewStoreInputFactory,
	NewUserInputFactory,
)

var repositorySet = wire.NewSet(
	NewStoreRepositoryFactory,
	NewUserRepositoryFactory,
)

var outputPortSet = wire.NewSet(
	NewStoreOutputFactory,
	NewUserOutputFactory,
)

var controllerSet = wire.NewSet(
	controller.NewStoreController,
	controller.NewUserController,
)

func InitializeRouter(ctx context.Context) (RouterI, error) {
	wire.Build(
		echoSet,
		driverSet,
		inputPortSet,
		repositorySet,
		outputPortSet,
		controllerSet,
		NewRouter,
	)
	return &Router{}, nil
}

func NewEcho() *echo.Echo {
	e := echo.New()
	e.Validator = NewValidator()
	return e
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
