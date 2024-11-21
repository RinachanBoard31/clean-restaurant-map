package gateway

import (
	"clean-storemap-api/src/driver/db"
	model "clean-storemap-api/src/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(*db.User) (*db.User, error) {
	args := m.Called()
	return args.Get(0).(*db.User), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(*db.User, map[string]interface{}) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) GetUser(int) (*db.User, error) {
	args := m.Called()
	return args.Get(0).(*db.User), args.Error(1)
}

func (m *MockUserRepository) FindByEmail(string) (*db.User, error) {
	args := m.Called()
	return args.Get(0).(*db.User), args.Error(1)
}
func (m *MockUserRepository) GenerateUrl() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *MockUserRepository) GetEmail(string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func TestCreate(t *testing.T) {
	/* Arrange */
	user := &model.User{
		Name:   "noiman",
		Email:  "noiman@groovex.co.jp",
		Age:    35,
		Sex:    1.0,
		Gender: -0.5,
	}
	dbUser := &db.User{
		Id:     1,
		Name:   user.Name,
		Email:  user.Email,
		Age:    user.Age,
		Sex:    user.Sex,
		Gender: user.Gender,
	}
	expected := user

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("CreateUser").Return(dbUser, nil)
	ug := &UserGateway{userDriver: mockUserRepository}

	/* Act */
	actual, _ := ug.Create(user)

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expected, actual)
	// userDriver.CreateUser()が1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "CreateUser", 1)
}

func TestExist(t *testing.T) {
	/* Arrange */
	user := &model.User{
		Name:   "noiman",
		Email:  "noiman@groovex.co.jp",
		Age:    35,
		Sex:    1.0,
		Gender: -0.5,
	}
	var expected error = nil
	foundUser := &db.User{}

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("FindByEmail").Return(foundUser, nil)
	ug := &UserGateway{userDriver: mockUserRepository}

	/* Act */
	actual := ug.Exist(user)

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expected, actual)
	// userDriver.CreateUser()が1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "FindByEmail", 1)
	// mockUserRepository.AssertNumberOfCalls(t, "UpdateUser", 1)
}
func TestUpdate(t *testing.T) {
	/* Arrange */
	var expected error = nil
	// 更新されるUser
	user := &model.User{
		Id:     1,
		Name:   "sample",
		Email:  "sample@example.com",
		Age:    10,
		Sex:    0.1,
		Gender: -0.1,
	}
	// 更新後のデータ
	updatedUserData := map[string]interface{}{"name": "sample2", "sex": 1.0, "gender": -1.0}

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("UpdateUser").Return(nil)
	ug := &UserGateway{userDriver: mockUserRepository}
	// ug := &UserGateway{userDriver: &db.DbUserDriver{}}

	/* Act */
	actual := ug.Update(user, updatedUserData)

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expected, actual)
	// userDriver.CreateUser()が1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "UpdateUser", 1)
}

func TestGet(t *testing.T) {
	/* Arrange */
	id := 1
	var expectedError error = nil
	dbUser := &db.User{
		Id:     id,
		Name:   "sample",
		Email:  "sample@example.com",
		Age:    20,
		Sex:    1.0,
		Gender: -0.5,
	}

	expectedUser := &model.User{
		Id:     dbUser.Id,
		Name:   dbUser.Name,
		Email:  dbUser.Email,
		Age:    dbUser.Age,
		Sex:    dbUser.Sex,
		Gender: dbUser.Gender,
	}

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GetUser").Return(dbUser, nil)
	ug := &UserGateway{userDriver: mockUserRepository}

	/* Act */
	actual, actualErr := ug.Get(id)

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expectedUser, actual)
	assert.Equal(t, expectedError, actualErr)
	// userDriver.CreateUser()が1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "GetUser", 1)
}

func TestFindBy(t *testing.T) {
	/* Arrange */
	var expected error = nil
	user := &model.UserCredentials{
		Email: "noiman@groovex.co.jp",
	}
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("FindByEmail").Return(&db.User{}, nil)
	ug := &UserGateway{userDriver: mockUserRepository}

	/* Act */
	actual := ug.FindBy(user)

	/* Assert */
	// 返り値が正しいこと
	assert.Equal(t, expected, actual)
	// userDriver.FindByEmail()が1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "FindByEmail", 1)
}

func TestGenerateAuthUrl(t *testing.T) {
	/* Arrange */
	expected := "https://www.google.com"
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GenerateUrl").Return(expected)
	ug := &UserGateway{
		googleOAuthDriver: mockUserRepository,
	}

	/* Act */
	actual := ug.GenerateAuthUrl()

	/* Assert */
	assert.Equal(t, expected, actual)
	mockUserRepository.AssertNumberOfCalls(t, "GenerateUrl", 1)
}

func TestGetUserInfoWithAuthCode(t *testing.T) {
	/* Arrange */
	code := ""
	email := "sample@example.com"
	expected := email
	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GetEmail").Return(email, nil)
	ug := &UserGateway{
		googleOAuthDriver: mockUserRepository,
	}

	/* Act */
	actual, _ := ug.GetUserInfoWithAuthCode(code)

	/* Assert */
	assert.Equal(t, expected, actual)
	mockUserRepository.AssertNumberOfCalls(t, "GetEmail", 1)
}
