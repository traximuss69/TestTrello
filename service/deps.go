package service

import "awesomeProject2/cmd/model"

type BoardStorage interface {
	GetBoards() []model.Board
	CreateBoard(title string) model.Board
}

type ListStorage interface {
	GetLists(boardID *int) []model.List
	CreateList(title string, boardID int) model.List
}

type CardStorage interface {
	GetCards(boardID *int) []model.Card
	CreateCard(title string, boardID int, listID int, description string) model.Card
	DeleteCard(boardID int, listID int, cardID int) (model.Card, error)
	UpdateCard(updated model.Card) (model.Card, error)
}
