package service

import "awesomeProject2/cmd/model"

type BoardService struct {
	storage BoardStorage
}

func NewBoardService(storage BoardStorage) *BoardService {
	return &BoardService{storage: storage}
}
func (s BoardService) GetBoards(BoardID *int) []model.Board {
	return s.storage.GetBoards(BoardID)
}
func (s BoardService) CreateBoard(title string) model.Board {
	return s.storage.CreateBoard(title)
}
