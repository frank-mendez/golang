package services

import (
	"golang/mvc/domain"
	"golang/mvc/utils"
)

type userService struct{}

var (
	UserService userService
)

func (u *userService) GetUsers(userId int64) (*domain.User,
	*utils.ApplicationError) {
	user, err := domain.UserDao.GetUser(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
