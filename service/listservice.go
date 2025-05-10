package service

import "awesomeProject2/cmd/model"

type ListService struct {
	storage ListStorage
}

func NewListService(storage ListStorage) *ListService {
	return &ListService{storage: storage}
}
func (s ListService) GetLists(ListID *int) []model.List {
	return s.storage.GetLists(ListID)
}
func (s ListService) CreateList(title string, boardID int) model.List {
	return s.storage.CreateList(title, boardID)
}
