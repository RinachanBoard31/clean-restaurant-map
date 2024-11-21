package interactor

import (
	model "clean-storemap-api/src/entity"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

type MockUserOutputPort struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) (*model.User, error) {
	args := m.Called()
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) Exist(user *model.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) GenerateAuthUrl() string {
	args := m.Called()
	return args.Get(0).(string)
}

func (m *MockUserRepository) FindBy(user *model.UserCredentials) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) GetUserInfoWithAuthCode(string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func (m *MockUserOutputPort) OutputCreateResult() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserOutputPort) OutputLoginResult() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserOutputPort) OutputAuthUrl(url string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserOutputPort) OutputSignupWithAuth(id int) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserOutputPort) OutputAlreadySignedup() error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	/* Arrange */
	var expected error = nil
	user := &model.User{Name: "natori", Email: "test@example.com", Age: 52, Sex: -0.2, Gender: 1.0}
	returnedUser := user
	returnedUser.Id = 1

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("Create").Return(returnedUser, nil)
	mockUserOutputPort := new(MockUserOutputPort)
	mockUserOutputPort.On("OutputCreateResult").Return(nil)

	ui := &UserInteractor{userRepository: mockUserRepository, userOutputPort: mockUserOutputPort}

	/* Act */
	actual := ui.CreateUser(user)

	/* Assert */
	assert.Equal(t, expected, actual)
	mockUserRepository.AssertNumberOfCalls(t, "Create", 1)
	mockUserOutputPort.AssertNumberOfCalls(t, "OutputCreateResult", 1)
}

func TestLoginUser(t *testing.T) {
	/* Arrange */
	var expected error = nil
	user := &model.UserCredentials{Email: "test@example.com"}

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("FindBy").Return(nil)
	mockUserOutputPort := new(MockUserOutputPort)
	mockUserOutputPort.On("OutputLoginResult").Return(nil)

	ui := &UserInteractor{userRepository: mockUserRepository, userOutputPort: mockUserOutputPort}

	/* Act */
	actual := ui.LoginUser(user)

	/* Assert */
	// LoginUser()がOutputLoginResult()を返すこと
	assert.Equal(t, expected, actual)
	// RepositoryのFindBy()が1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "FindBy", 1)
	// OutputPortのOutputLoginResult()が1回呼ばれること
	mockUserOutputPort.AssertNumberOfCalls(t, "OutputLoginResult", 1)
}

func TestGetAuthUrl(t *testing.T) {
	/* Arrange */
	url := "https://www.google.com"
	var expected error = nil

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GenerateAuthUrl").Return(url)
	mockUserOutputPort := new(MockUserOutputPort)
	mockUserOutputPort.On("OutputAuthUrl").Return(nil)

	ui := &UserInteractor{
		userRepository: mockUserRepository,
		userOutputPort: mockUserOutputPort,
	}

	/* Act */
	actual := ui.GetAuthUrl()

	/* Assert */
	assert.Equal(t, expected, actual)
	mockUserRepository.AssertNumberOfCalls(t, "GenerateAuthUrl", 1)
	mockUserOutputPort.AssertNumberOfCalls(t, "OutputAuthUrl", 1)
}

func TestSignupDraft(t *testing.T) {
	/* Arrange */
	code := ""
	email := "sample@example.com"
	var expected error = nil
	err := errors.New("user is not found")

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("GetUserInfoWithAuthCode").Return(email, nil)
	mockUserRepository.On("Exist").Return(err) // 存在していない場合にエラーが返る
	mockUserRepository.On("Create").Return(&model.User{}, nil)
	mockUserOutputPort := new(MockUserOutputPort)
	mockUserOutputPort.On("OutputAlreadySignedup").Return(nil)
	mockUserOutputPort.On("OutputSignupWithAuth").Return(nil)

	ui := &UserInteractor{
		userRepository: mockUserRepository,
		userOutputPort: mockUserOutputPort,
	}

	/* Act */
	actual := ui.SignupDraft(code)

	/* Assert */
	assert.Equal(t, expected, actual)
	mockUserRepository.AssertNumberOfCalls(t, "GetUserInfoWithAuthCode", 1)
	mockUserRepository.AssertNumberOfCalls(t, "Exist", 1)
	mockUserRepository.AssertNumberOfCalls(t, "Create", 1)
	mockUserOutputPort.AssertNumberOfCalls(t, "OutputSignupWithAuth", 1)
}
