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
	CreateUser(*db.User) (*db.User, error)
	UpdateUser(*db.User, map[string]interface{}) error
	GetUser(int) (*db.User, error)
	FindByEmail(string) (*db.User, error)
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

func (ug *UserGateway) Create(user *model.User) (*model.User, error) {
	dbUser := &db.User{
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Sex:    user.Sex,
		Gender: user.Gender,
	}

	dbUser, err := ug.userDriver.CreateUser(dbUser)
	if err != nil {
		return nil, err
	}
	user.Id = dbUser.Id // createが成功していればidを取得できるのでセットする
	return user, nil
}

func (ug *UserGateway) Exist(user *model.User) error {
	if _, err := ug.userDriver.FindByEmail(user.Email); err != nil {
		return err
	}
	return nil
}

func (ug *UserGateway) Update(user *model.User, updatedUserData map[string]interface{}) error {
	// updateされるUserをdb.Userに変換
	dbUser := &db.User{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Sex:    user.Sex,
		Gender: user.Gender,
	}
	if err := ug.userDriver.UpdateUser(dbUser, updatedUserData); err != nil {
		return err
	}
	return nil
}

func (ug *UserGateway) Get(id int) (*model.User, error) {
	dbUser, err := ug.userDriver.GetUser(id)
	if err != nil {
		return nil, err
	}
	user := &model.User{
		Id:     dbUser.Id,
		Name:   dbUser.Name,
		Email:  dbUser.Email,
		Age:    dbUser.Age,
		Sex:    dbUser.Sex,
		Gender: dbUser.Gender,
	}
	return user, nil
}

func (ug *UserGateway) FindBy(user *model.UserCredentials) error {
	if _, err := ug.userDriver.FindByEmail(user.Email); err != nil {
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
