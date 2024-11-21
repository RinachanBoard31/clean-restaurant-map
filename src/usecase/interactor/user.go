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

	if _, err := ui.userRepository.Create(user); err != nil {
		return err
	}
	if err := ui.userOutputPort.OutputCreateResult(); err != nil {
		return err
	}
	return nil
}

func (ui *UserInteractor) UpdateUser(updatedData map[string]interface{}) error {
	// データの整形(entityのカラムの有効範囲に基づいて整形するので、controllerではなくusecaseで行う)
	formatedData, err := model.UserFormat(updatedData)
	if err != nil {
		return err
	}
	// idを取得しidは更新しないので削除
	id := formatedData["id"].(int)
	delete(formatedData, "id")

	// userが存在するか確認
	user, err := ui.userRepository.Get(id)
	if err != nil {
		return err
	}

	if err := ui.userRepository.Update(user, formatedData); err != nil {
		return err
	}
	if err := ui.userOutputPort.OutputUpdateResult(); err != nil {
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
	// 存在しない場合にerrが返ってくるため、nilであればすでに存在しているということ
	if err := ui.userRepository.Exist(user); err == nil {
		// すでに登録されている場合はログイン画面に遷移させる
		if err := ui.userOutputPort.OutputAlreadySignedup(); err != nil {
			return err
		}
		return err
	}
	if user, err = ui.userRepository.Create(user); err != nil {
		return err
	}
	// urlのクエリパラメータにidを付与してそのidをユーザの更新時に受け取りどのユーザを更新するかを判別する
	if err := ui.userOutputPort.OutputSignupWithAuth(user.Id); err != nil {
		return err
	}
	return nil
}
