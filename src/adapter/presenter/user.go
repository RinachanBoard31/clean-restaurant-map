package presenter

import (
	"clean-storemap-api/src/usecase/port"
	"net/http"

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
