package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-playground/validator.v9"

	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
)

type UserController struct {
	userInputPort port.UserInputPort
}

type UserRequestBody struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (uc UserController) CreateUser(c echo.Context) error {
	var u UserRequestBody
	if err := c.Bind(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	if err := c.Validate(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, err.(validator.ValidationErrors).Error())
	}
	user, err := model.NewUser(u.Name, u.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return uc.userInputPort.CreateUser(user)
}
