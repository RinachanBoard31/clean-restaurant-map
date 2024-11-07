package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
	LoginUser(*model.UserCredentials) error
	GetGoogleAuthUrl() string
}

type UserRepository interface {
	Create(*model.User) error
	FindBy(*model.UserCredentials) error
	GenerateGoogleAuthUrl() string
}

type UserOutputPort interface {
	OutputCreateResult() error
	OutputLoginResult() error
	OutputGoogleAuthUrl(url string) string
}
