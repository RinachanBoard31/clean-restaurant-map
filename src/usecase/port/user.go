package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
	LoginUser(*model.UserCredentials) error
	GetAuthUrl() error
	SignupDraft(string) error
}

type UserRepository interface {
	Create(*model.User) (*model.User, error)
	FindBy(*model.UserCredentials) error
	GenerateAuthUrl() string
	GetUserInfoWithAuthCode(string) (string, error)
}

type UserOutputPort interface {
	OutputCreateResult() error
	OutputLoginResult() error
	OutputAuthUrl(string) error
	OutputSignupWithAuth() error
}
