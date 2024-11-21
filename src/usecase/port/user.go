package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
	UpdateUser(map[string]interface{}) error
	LoginUser(*model.UserCredentials) error
	GetAuthUrl() error
	SignupDraft(string) error
}

type UserRepository interface {
	Exist(*model.User) error
	Create(*model.User) (*model.User, error)
	Update(*model.User, map[string]interface{}) error
	Get(int) (*model.User, error)
	FindBy(*model.UserCredentials) error
	GenerateAuthUrl() string
	GetUserInfoWithAuthCode(string) (string, error)
}

type UserOutputPort interface {
	OutputCreateResult() error
	OutputUpdateResult() error
	OutputLoginResult() error
	OutputAuthUrl(string) error
	OutputSignupWithAuth(int) error
	OutputAlreadySignedup() error
}
