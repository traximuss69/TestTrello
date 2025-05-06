package service

import (
	"awesomeProject2/cmd/model"
)

type BoardService struct {
	storage BoardStorage
}
type ListService struct {
	storage ListStorage
}
type CardService struct {
	storage CardStorage
}

func NewService(storage BoardService) *BoardService {
	return &BoardService{storage: storage}
}
func (s BoardService) GetBoards(BoardID *int) []model.Board {
	return s.storage.GetBoards(BoardID)
}
func (s BoardService) CreateBoard(title string) model.Board {
	return s.storage.CreateBoard(title)
}
func (s ListService) GetList(ListID *int) []model.List {
	return s.storage.GetLists(ListID)
}
func (s ListService) CreateList(boardID int, title string) model.List {
	return s.storage.CreateList(title, boardID)
}
func (s CardService) GetCard(CardID *int) []model.Card {
	return s.storage.GetCards(CardID)
}
func (s CardService) CreateCard(boardID int, listID int, title string, description string) model.Card {
	return s.storage.CreateCard(title, boardID, listID, description)
}
func (s CardService) DeleteCard(boardID int, listID int, cardID int) (model.Card, error) {
	return s.storage.DeleteCard(boardID, listID, cardID)
}
func (s CardService) UpdatedCard(updated model.Card) (model.Card, error) {
	return s.storage.UpdateCard(updated)
}
