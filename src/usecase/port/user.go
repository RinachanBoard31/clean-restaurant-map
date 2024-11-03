package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
	LoginUser(*model.UserCredentials) error
}

type UserRepository interface {
	Create(*model.User) error
	FindBy(*model.UserCredentials) error
}

type UserOutputPort interface {
	OutputCreateResult() error
	OutputLoginResult() error
}
