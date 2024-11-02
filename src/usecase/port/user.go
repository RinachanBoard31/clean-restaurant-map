package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
	GetGoogleAuthUrl() string
}

type UserRepository interface {
	Create(*model.User) error
	GenerateGoogleAuthUrl() string
}

type UserOutputPort interface {
	OutputCreateResult() error
	OutputGoogleAuthUrl(url string) string
}
