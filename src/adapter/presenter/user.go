package presenter

import (
	"clean-storemap-api/src/usecase/port"
	"net/http"
	"os"
	"time"

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

func (up *UserPresenter) OutputLoginResult(token string) error {
	up.c.SetCookie(settingCookie(token))
	return up.c.JSON(http.StatusOK, map[string]interface{}{})
}

func (up *UserPresenter) OutputAuthUrl(url string) error {
	return up.c.Redirect(http.StatusFound, url)
}

func (up *UserPresenter) OutputSignupWithAuth(token string) error {
	url := os.Getenv("FRONT_URL") + "/editUser" // 認証以外のユーザ情報を入力するページ
	up.c.SetCookie(settingCookie(token))
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

func settingCookie(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = os.Getenv("JWT_TOKEN_NAME")
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour) // 24時間有効
	cookie.Path = "/"
	cookie.SameSite = http.SameSiteNoneMode // クロスサイトリクエストを許可
	cookie.HttpOnly = true
	cookie.Secure = true
	return cookie
}
