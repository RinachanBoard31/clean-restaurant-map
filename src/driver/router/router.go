package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"context"

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

type RouterI interface {
	Serve(ctx context.Context)
}

type Router struct {
	echo            *echo.Echo
	storeController controller.StoreI
	userController  controller.UserI
}

func NewRouter(echo *echo.Echo, storeController controller.StoreI, userController controller.UserI) RouterI {
	return &Router{
		echo:            echo,
		storeController: storeController,
		userController:  userController,
	}
}

func (router *Router) Serve(ctx context.Context) {
	router.echo.GET("/", router.storeController.GetStores)
	router.echo.POST("/user", router.userController.CreateUser)
	router.echo.GET("/stores/opening-hours", router.storeController.GetNearStores)
	router.echo.GET("/auth", router.userController.GetGoogleAuthUrl) // Google認証用のURLを取得し返す(ユーザの登録はsignupで行う)
	// router.echo.POST("/signup", router.userController.SignUp) // Google認証後のユーザー登録(未実装)
	router.echo.Logger.Fatal(router.echo.Start(":8080"))
}
