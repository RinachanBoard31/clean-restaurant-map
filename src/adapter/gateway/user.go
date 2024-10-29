package gateway

import (
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
	"errors"
)

type UserGateway struct {
	userDriver UserDriver
}

type UserDriver interface {
	CreateUser(*db.User) error
	CheckUser(string) error
}

func NewUserRepository(userDriver UserDriver) port.UserRepository {
	return &UserGateway{
		userDriver: userDriver,
	}
}

func (ug *UserGateway) Create(user *model.User) error {
	dbUser := &db.User{
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Sex:    user.Sex,
		Gender: user.Gender,
	}
	if err := ug.userDriver.CreateUser(dbUser); err != nil {
		return err
	}
	return nil
}

func (ug *UserGateway) Check(user *model.UserCredentials) error {
	if err := ug.userDriver.CheckUser(user.Email); err != nil {
		return errors.New("emailが異なっています。")
	}
	return nil
}
