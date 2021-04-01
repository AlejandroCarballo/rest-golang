package service

import "github.com/AlejandroCarballo/ejemplo/internal/models"

// ItemService ...
type ItemService interface {
	Get(int) models.Item
	Describe() string
}

type itemService struct {
}

// NewItemService ...
func NewItemService() ItemService {
	return &itemService{}
}

func (is *itemService) Describe() string {
	return "Hello World"
}

func (is *itemService) Get(ID int) models.Item {
	return models.Item{}
}
