package storage

import (
	"awesomeProject2/cmd/model"
	"time"
)

type Storage struct {
	Boards  []model.Board
	boardID int
	listID  int
	cardID  int
}

func NewStorage() *Storage {
	return &Storage{
		Boards:  []model.Board{},
		boardID: 1,
		listID:  1,
		cardID:  1,
	}
}
func (s *Storage) GetBoards(BoardID *int) []model.Board {
	var result []model.Board
	for _, b := range s.Boards {
		if BoardID != nil {
			if b.ID == *BoardID {
				return []model.Board{b}
			}
		} else {
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
	board.ID = s.boardID
	s.boardID++
	s.Boards = append(s.Boards, board)
	return board
}
func (s *Storage) GetList(ListID *int) []model.List {
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
func (s *Storage) CreateList(boardID int, title string) model.List {
	newList := model.List{
		Title:     title,
		BoardID:   boardID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	newList.ID = s.listID
	s.listID++
	newList.CreatedAt = time.Now()
	newList.UpdatedAt = time.Now()
	for i := range s.Boards {
		if s.Boards[i].ID == boardID {
			s.Boards[i].Lists = append(s.Boards[i].Lists, newList)
			return newList
		}
	}
	return model.List{}
}
func (s *Storage) GetCard(CardID *int) []model.Card {
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

func (s *Storage) CreateCard(boardID int, listID int, title string, description string) model.Card {
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
func (s *Storage) DeleteCard(boardID int, listID int, cardID int) (model.Card, bool) {
	for i := range s.Boards {
		if s.Boards[i].ID == boardID {
			for j := range s.Boards[i].Lists {
				if s.Boards[i].Lists[j].ID == listID {
					cards := s.Boards[i].Lists[j].Cards
					for k, c := range cards {
						if c.ID == cardID {
							s.Boards[i].Lists[j].Cards = append(cards[:k], cards[k+1:]...)
							return c, true
						}
					}
				}
			}
		}
	}
	return model.Card{}, false
}
func (s *Storage) UpdatedCard(updated model.Card) (model.Card, bool) {
	for i := range s.Boards {
		if s.Boards[i].ID == updated.BoardID {
			for j := range s.Boards[i].Lists {
				if s.Boards[i].Lists[j].ID == updated.ListID {
					for c := range s.Boards[i].Lists[j].Cards {
						if s.Boards[i].Lists[j].Cards[c].ID == updated.ID {
							s.Boards[i].Lists[j].Cards[c].Title = updated.Title
							s.Boards[i].Lists[j].Cards[c].Description = updated.Description
							return s.Boards[i].Lists[j].Cards[c], true
						}
					}
				}
			}
		}
	}
	return model.Card{}, false
}
