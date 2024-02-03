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
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
