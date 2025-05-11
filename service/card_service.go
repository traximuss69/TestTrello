package service

import "awesomeProject2/cmd/model"

type CardService struct {
	storage CardStorage
}

func NewCardService(storage CardStorage) *CardService {
	return &CardService{storage: storage}
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
