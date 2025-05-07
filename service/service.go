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

func NewBoardService(storage BoardStorage) *BoardService {
	return &BoardService{storage: storage}
}
func NewListService(storage ListStorage) *ListService {
	return &ListService{storage: storage}
}
func NewCardService(storage CardStorage) *CardService {
	return &CardService{storage: storage}
}

func (s BoardService) GetBoards(BoardID *int) []model.Board {
	return s.storage.GetBoards(BoardID)
}
func (s BoardService) CreateBoard(title string) model.Board {
	return s.storage.CreateBoard(title)
}
func (s ListService) GetLists(ListID *int) []model.List {
	return s.storage.GetLists(ListID)
}
func (s ListService) CreateList(title string, boardID int) model.List {
	return s.storage.CreateList(title, boardID)
}
func (s CardService) GetCards(CardID *int) []model.Card {
	return s.storage.GetCards(CardID)
}
func (s CardService) CreateCard(title string, boardID int, listID int, description string) model.Card {
	return s.storage.CreateCard(title, boardID, listID, description)
}
func (s CardService) DeleteCard(boardID int, listID int, cardID int) (model.Card, error) {
	return s.storage.DeleteCard(boardID, listID, cardID)
}
func (s CardService) UpdateCard(updated model.Card) (model.Card, error) {
	return s.storage.UpdateCard(updated)
}
