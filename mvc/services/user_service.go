package services

import "golang/mvc/domain"

func GetUsers(userId int64) (*domain.User,
	error) {
	return domain.GetUser(userId)
}
