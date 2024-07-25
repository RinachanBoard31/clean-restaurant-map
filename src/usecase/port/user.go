package port

import (
	model "clean-storemap-api/src/entity"
)

type UserInputPort interface {
	CreateUser(*model.User) error
}
