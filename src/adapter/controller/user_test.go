package controller

import (
	"bytes"
	model "clean-storemap-api/src/entity"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserInputPort struct {
	mock.Mock
}

func (m *MockUserInputPort) CreateUser(*model.User) error {
	args := m.Called()
	return args.Error(0)
}

func TestCreateUser(t *testing.T) {
	/* Arrange */
	c, rec := newRouter()
	expected := errors.New("")
	mockUserInputPort := new(MockUserInputPort)
	mockUserInputPort.On("CreateUser").Return(errors.New(""))
	uc := &UserController{userInputPort: mockUserInputPort}

	reqBody := `{"name":"test","email":"aaa@gmail.com"}`
	req := httptest.NewRequest(http.MethodPost, "/user", bytes.NewBufferString(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.SetRequest(req)

	/* Act */
	actual := uc.CreateUser(c)

	/* Assert */
	// uc.CreateUser()がUserInputPort.CreateUser()を返すこと
	assert.Equal(t, expected, actual)
	// echoが正しく起動したか
	assert.Equal(t, http.StatusOK, rec.Code)
	// InputPortのCreateUserが1回呼ばれること
	mockUserInputPort.AssertNumberOfCalls(t, "CreateUser", 1)
}
