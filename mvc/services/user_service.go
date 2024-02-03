package services

import (
	"golang/mvc/domain"
	"golang/mvc/utils"
)

func GetUsers(userId int64) (*domain.User,
	*utils.ApplicationError) {
	return domain.GetUser(userId)
}
