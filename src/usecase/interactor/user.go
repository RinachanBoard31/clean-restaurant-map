package interactor

import (
	model "clean-storemap-api/src/entity"
	port "clean-storemap-api/src/usecase/port"
)

type UserInteractor struct {
	userRepository port.UserRepository
	userOutputPort port.UserOutputPort
}

func NewUserInputPort(userRepository port.UserRepository, userOutputPort port.UserOutputPort) port.UserInputPort {
	return &UserInteractor{
		userRepository: userRepository,
		userOutputPort: userOutputPort,
	}
}

func (ui *UserInteractor) CreateUser(user *model.User) error {
	if err := ui.userRepository.Create(user); err != nil {
		return err
	}
	if err := ui.userOutputPort.OutputCreateResult(); err != nil {
		return err
	}
	return nil
}

func (ui *UserInteractor) LoginUser(user *model.UserCredentials) error {
	if err := ui.userRepository.FindBy(user); err != nil {
		return err
	}
	if err := ui.userOutputPort.OutputLoginResult(); err != nil {
		return err
	}
	return nil
}

func (ui *UserInteractor) GetGoogleAuthUrl() string {
	url := ui.userRepository.GenerateGoogleAuthUrl()
	return ui.userOutputPort.OutputGoogleAuthUrl(url)
}
