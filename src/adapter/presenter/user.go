package presenter

import (
	"clean-storemap-api/src/usecase/port"
	"net/http"
	"os"
	"strconv"

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

func (up *UserPresenter) OutputLoginResult() error {
	return up.c.JSON(http.StatusOK, map[string]interface{}{})
}

func (up *UserPresenter) OutputAuthUrl(url string) error {
	return up.c.Redirect(http.StatusFound, url)
}

func (up *UserPresenter) OutputSignupWithAuth(id int) error {
	url := os.Getenv("FRONT_URL") + "/editUser" + "?id=" + strconv.Itoa(id) // 認証以外のユーザ情報を入力するページ
	return up.c.Redirect(http.StatusFound, url)
}
