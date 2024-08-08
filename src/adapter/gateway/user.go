package gateway

import (
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
)

type UserGateway struct {
	userDriver UserDriver
}

type UserDriver interface {
	CreateUser(*db.User) error
}

func (ug *UserGateway) Create(user *model.User) error {
	dbUser := &db.User{
		Name:  user.Name,
		Email: user.Email,
	}
	if err := ug.userDriver.CreateUser(dbUser); err != nil {
		return err
	}
	return nil
}
