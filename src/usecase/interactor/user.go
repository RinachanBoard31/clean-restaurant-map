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

func (ui *UserInteractor) GetAuthUrl() error {
	url := ui.userRepository.GenerateAuthUrl()
	return ui.userOutputPort.OutputAuthUrl(url)
}

func (ui *UserInteractor) SignupDraft(code string) error {
	email, err := ui.userRepository.GetUserInfoWithAuthCode(code)
	if err != nil {
		return err
	}

	// 先にemailのみで登録する(仮登録)
	user := &model.User{
		Name:   "",
		Email:  email,
		Age:    0,
		Sex:    0.0,
		Gender: 0.0,
	}
	if err := ui.userRepository.Create(user); err != nil {
		return err
	}

	if err := ui.userOutputPort.OutputSignupWithAuth(); err != nil {
		return err
	}
	return nil
}
