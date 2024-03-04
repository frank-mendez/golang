package services

import (
	"github.com/stretchr/testify/assert"
	"golang/mvc/domain"
	"golang/mvc/utils"
	"net/http"
	"testing"
)

var (
	userDaoMock     usersDaoMock
	getUserFunction func(userId int64) (*domain.User, *utils.ApplicationError)
)

func init() {
	domain.UserDao = &usersDaoMock{}
}

type usersDaoMock struct{}

func (m *usersDaoMock) GetUser(userId int64) (*domain.User, *utils.ApplicationError) {
	return getUserFunction(userId)
}

func TestGetUserNotFoundInDatabase(t *testing.T) {
	// Initialization:
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return nil, &utils.ApplicationError{StatusCode: http.StatusNotFound, Message: "user 0 was not found"}
	}
	user, err := UserService.GetUsers(0)
	// Execution:

	// Validation:
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, "user 0 was not found", err.Message)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
}

func TestGetUserNoError(t *testing.T) {
	// Initialization:
	getUserFunction = func(userId int64) (*domain.User, *utils.ApplicationError) {
		return &domain.User{Id: 123}, nil
	}
	user, err := UserService.GetUsers(123)
	// Execution:

	// Validation:
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.EqualValues(t, 123, user.Id)
}
