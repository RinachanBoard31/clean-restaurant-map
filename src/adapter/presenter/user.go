package presenter

import (
	"clean-storemap-api/src/usecase/port"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type UserPresenter struct {
	c echo.Context
}

func NewUserOutputPort(c echo.Context) port.UserOutputPort {
	return &UserPresenter{c: c}
}

func (up *UserPresenter) OutputCreateResult() error {
	return up.c.JSON(http.StatusOK, map[string]interface{}{})
}

func (up *UserPresenter) OutputUpdateResult() error {
	return up.c.JSON(http.StatusOK, map[string]interface{}{})
}

func (up *UserPresenter) OutputLoginResult(userId string) error {
	return up.c.JSON(http.StatusOK, map[string]interface{}{"userId": userId})
}

func (up *UserPresenter) OutputAuthUrl(url string) error {
	return up.c.Redirect(http.StatusFound, url)
}

func (up *UserPresenter) OutputSignupWithAuth(id string) error {
	url := os.Getenv("FRONT_URL") + "/editUser" + "?id=" + id // 認証以外のユーザ情報を入力するページ
	return up.c.Redirect(http.StatusFound, url)
}

func (up *UserPresenter) OutputAlreadySignedup() error {
	url := os.Getenv("FRONT_URL") // すでに登録済みの場合はトップページにリダイレクト
	return up.c.Redirect(http.StatusFound, url)
}

func (up *UserPresenter) OutputHasEmailInRequestBody() error {
	errMsg := "Email is included in Request Body"
	return up.c.JSON(http.StatusBadRequest, map[string]interface{}{"error": errMsg})
}
