package router

import (
	controller "clean-storemap-api/src/adapter/controller"
	"clean-storemap-api/src/driver/middleware"

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
	// ログイン前のルーティング
	router.echo.GET("/", router.storeController.GetStores)
	router.echo.POST("/user", router.userController.CreateUser) // こっちは時期に使わなくなります。
	router.echo.POST("/login", router.userController.LoginUser)
	router.echo.GET("/auth", router.userController.GetAuthUrl)            // Google認証用のURLを取得し返す
	router.echo.GET("/auth/signup", router.userController.SignupWithAuth) // ユーザの認証を確認し仮登録する(本登録は未実装,UpdateUserで行う)

	// ログイン後のルーティング
	afterLoginPath := router.echo.Group("")
	middleware.SetupJwtMiddleware(afterLoginPath)

	afterLoginPath.GET("/stores/opening-hours", router.storeController.GetNearStores)
	afterLoginPath.GET("/stores/favorite-ranking", router.storeController.GetTopFavoriteStores)
	afterLoginPath.GET("/user/:user_id/favorite-store", router.storeController.GetFavoriteStores)
	afterLoginPath.POST("/user/:user_id/favorite-store", router.storeController.SaveFavoriteStore)
	afterLoginPath.PUT("/user/:id", router.userController.UpdateUser)
	router.echo.Logger.Fatal(router.echo.Start(":8080"))
}
