package domain

import (
	"fmt"
	"golang/mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		23: {Id: 23, FirstName: "Frank", LastName: "Mendez", Email: "frankmendezwebdev@gmail.com"},
	}
	UserDao usersDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
