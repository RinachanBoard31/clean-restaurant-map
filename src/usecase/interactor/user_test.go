package interactor

import (
	model "clean-storemap-api/src/entity"
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

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserOutputPort) OutputCreateResult() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserRepository) Check(user *model.UserCredentials) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockUserOutputPort) OutputCheckResult() error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	/* Arrange */
	var expected error = nil
	user := &model.User{Id: 1, Name: "natori", Email: "test@example.com", Age: 52, Sex: -0.2, Gender: 1.0}

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("Create").Return(nil)
	mockUserOutputPort := new(MockUserOutputPort)
	mockUserOutputPort.On("OutputCreateResult").Return(nil)

	ui := &UserInteractor{userRepository: mockUserRepository, userOutputPort: mockUserOutputPort}

	/* Act */
	actual := ui.CreateUser(user)

	/* Assert */
	// CreateUser()がOutputCreateResult()を返すこと
	assert.Equal(t, expected, actual)
	// RepositoryのCreateが1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "Create", 1)
	// OutputPortのOutputCreateResult()が1回呼ばれること
	mockUserOutputPort.AssertNumberOfCalls(t, "OutputCreateResult", 1)
}

func TestCheckUser(t *testing.T) {
	/* Arrange */
	var expected error = nil
	user := &model.UserCredentials{Email: "test@example.com"}

	mockUserRepository := new(MockUserRepository)
	mockUserRepository.On("Check").Return(nil)
	mockUserOutputPort := new(MockUserOutputPort)
	mockUserOutputPort.On("OutputCheckResult").Return(nil)

	ui := &UserInteractor{userRepository: mockUserRepository, userOutputPort: mockUserOutputPort}

	/* Act */
	actual := ui.CheckUser(user)

	/* Assert */
	// CheckUser()がOutputCheckResult()を返すこと
	assert.Equal(t, expected, actual)
	// RepositoryのCheckが1回呼ばれること
	mockUserRepository.AssertNumberOfCalls(t, "Check", 1)
	// OutputPortのOutputCheckResult()が1回呼ばれること
	mockUserOutputPort.AssertNumberOfCalls(t, "OutputCheckResult", 1)
}
