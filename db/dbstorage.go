package storage

import (
	"awesomeProject2/cmd/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sqlx.DB
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{DB: db}
}

func (s *Storage) GetBoards() []model.Board {
	var boards []model.Board
	err := s.DB.Select(&boards, "SELECT * FROM boards")
	if err != nil {
		return nil
	}
	return boards
}
func (s *Storage) CreateBoard(title string) model.Board {
	var board model.Board
	query := `INSERT INTO boards (title) VALUES ($1) RETURNING id, title`
	err := s.DB.Get(&board, query, title)
	if err != nil {
		return model.Board{}
	}
	return board
}
func (s *Storage) GetLists(boardID *int) []model.List {
	var lists []model.List
	err := s.DB.Select(&lists, "SELECT * FROM lists")
	if err != nil {
		return []model.List{}
	}
	return lists
}
func (s *Storage) CreateList(title string, boardID int) model.List {
	var list model.List
	query := `INSERT INTO lists (title, board_id) VALUES ($1, $2) RETURNING id, title, board_id`
	err := s.DB.Get(&list, query, title, boardID)
	if err != nil {
		return model.List{}
	}
	return list
}
func (s *Storage) GetCards(boardID *int) []model.Card {
	var cards []model.Card
	err := s.DB.Select(&cards, "SELECT * FROM cards")
	if err != nil {
		return []model.Card{}
	}
	return cards
}
func (s *Storage) CreateCard(title string, boardID int, listID int, description string) model.Card {
	query := `INSERT INTO cards(title, description, list_id) VALUES ($1, $2, $3) RETURNING id, title, description, list_id`
	var card model.Card
	err := s.DB.Get(&card, query, title, description, listID)
	if err != nil {
		return model.Card{}
	}
	return card
}
func (s *Storage) DeleteCard(boardID int, listID int, cardID int) (model.Card, error) {
	query := `DELETE FROM cards WHERE id = $1 AND board_id = $2 AND list_id = $3 RETURNING id, board_id, list_id`
	var card model.Card
	err := s.DB.Get(&card, query, cardID, boardID, listID)
	return card, err
}

func (s *Storage) UpdateCard(updated model.Card) (model.Card, error) {
	query := `UPDATE cards SET title = $1, description = $2, list_id = $3 WHERE id = $4 RETURNING id, title, description, list_id, board_id`
	var card model.Card
	err := s.DB.Get(&card, query, updated.Title, updated.Description, updated.ListID, updated.ID)
	return card, err
}
