package storage

import (
	"awesomeProject2/cmd/model"
	"fmt"
	"time"
)

type Storage struct {
	Boards  map[int]model.Board
	boardID int
	listID  int
	cardID  int
}

func NewStorage() *Storage {
	return &Storage{
		Boards: make(map[int]model.Board),
	}
}
func (s *Storage) GetBoards(boardID *int) []model.Board {
	var result []model.Board
	if boardID == nil {
		for _, b := range s.Boards {
			result = append(result, b)
		}
	} else {
		if b, ok := s.Boards[*boardID]; ok {
			result = append(result, b)
		}
	}
	return result
}
func (s *Storage) CreateBoard(title string) model.Board {
	board := model.Board{
		Title:     title,
		ID:        s.boardID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.Boards[s.boardID] = board
	s.boardID++
	return board
}
func (s *Storage) GetLists(ListID *int) []model.List {
	var result []model.List
	for _, b := range s.Boards {
		for _, l := range b.Lists {
			if ListID != nil {
				if l.ID == *ListID {
					return []model.List{l}
				}
			} else {
				result = append(result, l)
			}
		}
	}
	return result
}
func (s *Storage) CreateList(title string, boardID int) model.List {
	newList := model.List{
		ID:        s.listID,
		Title:     title,
		BoardID:   boardID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.listID++
	board, ok := s.Boards[boardID]
	if !ok {
		return model.List{}
	}
	board.Lists = append(board.Lists, newList)
	s.Boards[boardID] = board
	return newList
}
func (s *Storage) GetCards(CardID *int) []model.Card {
	var result []model.Card
	for _, b := range s.Boards {
		for _, l := range b.Lists {
			for _, c := range l.Cards {
				if CardID != nil {
					if c.ID == *CardID {
						return []model.Card{c}
					}
				} else {
					result = append(result, c)
				}
			}
		}
	}
	return result
}

func (s *Storage) CreateCard(title string, boardID int, listID int, description string) model.Card {
	newCard := model.Card{
		Title:       title,
		Description: description,
		ID:          s.cardID,
		BoardID:     boardID,
		ListID:      listID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newCard.ID = s.cardID
	s.cardID++
	newCard.CreatedAt = time.Now()
	newCard.UpdatedAt = time.Now()
	for i := range s.Boards {
		if s.Boards[i].ID == boardID {
			for j := range s.Boards[i].Lists {
				if s.Boards[i].Lists[j].ID == listID {
					s.Boards[i].Lists[j].Cards = append(s.Boards[i].Lists[j].Cards, newCard)
					return newCard
				}
			}
		}
	}
	return model.Card{}
}
func (s *Storage) DeleteCard(boardID int, listID int, cardID int) (model.Card, error) {
	for i := range s.Boards {
		if s.Boards[i].ID == boardID {
			for j := range s.Boards[i].Lists {
				if s.Boards[i].Lists[j].ID == listID {
					cards := s.Boards[i].Lists[j].Cards
					for k, c := range cards {
						if c.ID == cardID {
							s.Boards[i].Lists[j].Cards = append(cards[:k], cards[k+1:]...)
							return c, nil
						}
					}
					return model.Card{}, fmt.Errorf("card %d not found in list %d", cardID, listID)
				}
			}
			return model.Card{}, fmt.Errorf("list %d not found in board %d", listID, boardID)
		}
	}
	return model.Card{}, fmt.Errorf("board %d not found", boardID)
}
func (s *Storage) UpdateCard(updated model.Card) (model.Card, error) {
	for i := range s.Boards {
		if s.Boards[i].ID == updated.BoardID {
			for j := range s.Boards[i].Lists {
				if s.Boards[i].Lists[j].ID == updated.ListID {
					for c := range s.Boards[i].Lists[j].Cards {
						if s.Boards[i].Lists[j].Cards[c].ID == updated.ID {
							s.Boards[i].Lists[j].Cards[c].Title = updated.Title
							s.Boards[i].Lists[j].Cards[c].Description = updated.Description
							return s.Boards[i].Lists[j].Cards[c], nil
						}
					}
					return model.Card{}, fmt.Errorf("card %d not found in list %d", updated.ID, updated.ListID)
				}
			}
			return model.Card{}, fmt.Errorf("list %d not found in board %d", updated.ListID, updated.BoardID)
		}
	}
	return model.Card{}, fmt.Errorf("board %d not found", updated.BoardID)
}
