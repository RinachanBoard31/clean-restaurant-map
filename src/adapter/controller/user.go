package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	"clean-storemap-api/src/adapter/gateway"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
)

type UserI interface {
	CreateUser(c echo.Context) error
	LoginUser(c echo.Context) error
	GetGoogleAuthUrl(c echo.Context) error
}

type UserOutputFactory func(echo.Context) port.UserOutputPort
type UserInputFactory func(port.UserRepository, port.UserOutputPort) port.UserInputPort
type UserRepositoryFactory func(gateway.UserDriver, gateway.GoogleOAuthDriver) port.UserRepository
type UserDriverFactory gateway.UserDriver
type GoogleOAuthDriverFactory gateway.GoogleOAuthDriver

type UserController struct {
	userDriverFactory        UserDriverFactory
	googleOAuthDriverFactory GoogleOAuthDriverFactory
	userOutputFactory        UserOutputFactory
	userInputFactory         UserInputFactory
	userRepositoryFactory    UserRepositoryFactory
}

// 0が存在しないとして扱われるため数字型(int, float32)にvalidate:"required"を使用していない。(requiredがなくても型確認はされます。)
// 数字型のものが未入力であれば0として扱われる
// 0を存在する値とする場合にはカスタムバリデーションを使用する必要があり、カスタムバリデーションにはrouterで定義されたecho.New()を使用するため今回はカスタムバリデーションを使用しない。
type UserRequestBody struct {
	Name   string  `json:"name" validate:"required"`
	Email  string  `json:"email" validate:"required,email"`
	Age    int     `json:"age"`
	Sex    float32 `json:"sex"`
	Gender float32 `json:"gender"`
}
type UserCredentialsRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

func NewUserController(
	userDriverFactory UserDriverFactory,
	googleOAuthDriverFactory GoogleOAuthDriverFactory,
	userOutputFactory UserOutputFactory,
	userInputFactory UserInputFactory,
	userRepositoryFactory UserRepositoryFactory,
) UserI {
	return &UserController{
		userDriverFactory:        userDriverFactory,
		googleOAuthDriverFactory: googleOAuthDriverFactory,
		userOutputFactory:        userOutputFactory,
		userInputFactory:         userInputFactory,
		userRepositoryFactory:    userRepositoryFactory,
	}
}

func (uc *UserController) CreateUser(c echo.Context) error {
	var u UserRequestBody
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.(validator.ValidationErrors).Error())
	}
	user, err := model.NewUser(u.Name, u.Email, u.Age, u.Sex, u.Gender)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return uc.newUserInputPort(c).CreateUser(user)
}

func (uc *UserController) LoginUser(c echo.Context) error {
	// UserCredentialsの値を受け取りuserが既に登録されているかを確かめるキー用の型を作成する
	var u UserCredentialsRequestBody
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.(validator.ValidationErrors).Error())
	}
	user, err := model.NewUserCredentials(u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if err := uc.newUserInputPort(c).LoginUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return nil
}
func (uc *UserController) GetGoogleAuthUrl(c echo.Context) error {
	url := uc.newUserInputPort(c).GetGoogleAuthUrl()
	return c.Redirect(http.StatusFound, url) // 認証ページへのリダイレクトを行う
}

/* ここでpresenterにecho.Contextを渡している！起爆！！！（遅延） */
/* これによって、presenterのinterface(outputport)にecho.Contextを書かなくて良くなる */
func (uc *UserController) newUserInputPort(c echo.Context) port.UserInputPort {
	userOutputPort := uc.userOutputFactory(c)
	userDriver := uc.userDriverFactory
	googleOAuthDriver := uc.googleOAuthDriverFactory
	userRepository := uc.userRepositoryFactory(userDriver, googleOAuthDriver)
	return uc.userInputFactory(userRepository, userOutputPort)
}
