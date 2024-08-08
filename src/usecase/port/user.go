package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
}

type UserRepository interface {
	Create(*model.User) error
}

type UserOutputPort interface {
	OutputCreateResult() error
}