package gateway

import (
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"
)

type UserGateway struct {
	userDriver        UserDriver
	googleOAuthDriver GoogleOAuthDriver
}

type UserDriver interface {
	CreateUser(*db.User) error
	FindByEmail(string) error
}

type GoogleOAuthDriver interface {
	GenerateUrl() string
	GetEmail(string) (string, error)
}

func NewUserRepository(userDriver UserDriver, googleOAuthDriver GoogleOAuthDriver) port.UserRepository {
	return &UserGateway{
		userDriver:        userDriver,
		googleOAuthDriver: googleOAuthDriver,
	}
}

func (ug *UserGateway) Create(user *model.User) error {
	dbUser := &db.User{
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Sex:    user.Sex,
		Gender: user.Gender,
	}
	if err := ug.userDriver.CreateUser(dbUser); err != nil {
		return err
	}
	return nil
}

func (ug *UserGateway) FindBy(user *model.UserCredentials) error {
	if err := ug.userDriver.FindByEmail(user.Email); err != nil {
		return err
	}
	return nil
}
func (ug *UserGateway) GenerateAuthUrl() string {
	return ug.googleOAuthDriver.GenerateUrl()
}

func (ug *UserGateway) GetUserInfoWithAuthCode(code string) (string, error) {
	email, err := ug.googleOAuthDriver.GetEmail(code)
	if err != nil {
		return "", err
	}
	return email, nil
}
