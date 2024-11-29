package gateway

import (
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"clean-storemap-api/src/usecase/port"

	"github.com/google/uuid"
)

type UserGateway struct {
	userDriver        UserDriver
	googleOAuthDriver GoogleOAuthDriver
}

type UserDriver interface {
	CreateUser(*db.User) (*db.User, error)
	UpdateUser(*db.User, map[string]interface{}) error
	FindById(string) (*db.User, error)
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
		Id:     uuid.New().String(),
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

func (ug *UserGateway) Update(user *model.User, updateData model.ChangeForUser) error {
	// updateされるUserをdb.Userに変換
	dbUser := &db.User{
		Id:     user.Id,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Sex:    user.Sex,
		Gender: user.Gender,
	}
	if err := ug.userDriver.UpdateUser(dbUser, updateData); err != nil {
		return err
	}
	return nil
}

func (ug *UserGateway) Get(id string) (*model.User, error) {
	dbUser, err := ug.userDriver.FindById(id)
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

func (ug *UserGateway) FindBy(userCredentials *model.UserCredentials) (*model.User, error) {
	dbUser, err := ug.userDriver.FindByEmail(userCredentials.Email)
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
