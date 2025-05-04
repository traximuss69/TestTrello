package service

import (
	"awesomeProject2/cmd/model"
	"awesomeProject2/cmd/storage"
)

type Service struct {
	Storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{Storage: storage}
}
func (s Service) GetBoards(BoardID *int) []model.Board {
	return s.Storage.GetBoards(BoardID)
}
func (s Service) CreateBoard(title string) model.Board {
	return s.Storage.CreateBoard(title)
}
func (s Service) GetList(ListID *int) []model.List {
	return s.Storage.GetList(ListID)
}
func (s Service) CreateList(boardID int, title string) model.List {
	return s.Storage.CreateList(boardID, title)
}
func (s Service) GetCard(CardID *int) []model.Card {
	return s.Storage.GetCard(CardID)
}
func (s Service) CreateCard(boardID int, listID int, title string, description string) model.Card {
	return s.Storage.CreateCard(boardID, listID, title, description)
}
func (s Service) DeleteCard(boardID int, listID int, cardID int) (model.Card, bool) {
	return s.Storage.DeleteCard(boardID, listID, cardID)
}
func (s Service) UpdatedCard(updated model.Card) (model.Card, bool) {
	return s.Storage.UpdatedCard(updated)
}
