package service

import (
	"awesomeProject2/model"
	"awesomeProject2/storage"
	"encoding/json"
	"net/http"
)

type Service struct {
	Storage *storage.Storage
}

func NewService(storage *storage.Storage) *Service {
	return &Service{Storage: storage}
}
func (s Service) GetBoards() []string {
	return s.Storage.GetBoards()
}
func (s Service) CreateBoard(title string) model.Board {
	return s.Storage.CreateBoard(title)
}
func (s Service) GetList() []string {
	return s.Storage.GetList()
}
func (s Service) CreateList(boardID int, title string) model.List {
	return s.Storage.CreateList(boardID, title)
}
func (s Service) GetCard() []string {
	return s.Storage.GetCard()
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
func SendFullJSON(w http.ResponseWriter, board []model.Board) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
