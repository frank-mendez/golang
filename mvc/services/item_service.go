package services

import (
	"golang/mvc/domain"
	"golang/mvc/utils"
	"net/http"
)

type itemService struct{}

var (
	ItemService itemService
)

func (s *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "Implement me",
		StatusCode: http.StatusInternalServerError,
	}
}
