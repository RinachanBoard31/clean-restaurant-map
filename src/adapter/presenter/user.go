package presenter

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserPresenter struct {
	c echo.Context
}

func (up *UserPresenter) OutputCreateResult() error {
	return up.c.JSON(http.StatusOK, map[string]interface{}{})
}
