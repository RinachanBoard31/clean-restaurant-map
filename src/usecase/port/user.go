package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
	CheckUser(*model.UserCredentials) error
}

type UserRepository interface {
	Create(*model.User) error
	Check(*model.UserCredentials) error
}

type UserOutputPort interface {
	OutputCreateResult() error
	OutputCheckResult() error
}
