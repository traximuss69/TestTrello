package dto

import "awesomeProject2/cmd/model"

type BoardDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func BoardToDTO(b model.Board) BoardDTO {
	return BoardDTO{
		ID:    b.ID,
		Title: b.Title,
	}
}

type ListDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	BoardID int    `json:"board_id"`
}

func ListToDTO(l model.List) ListDTO {
	return ListDTO{
		ID:      l.ID,
		Title:   l.Title,
		BoardID: l.BoardID,
	}
}

type CardDTO struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ListID      int    `json:"list_id"`
}
type UpdatedCardDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func CardToDTO(c model.Card) CardDTO {
	return CardDTO{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		ListID:      c.ListID,
	}
}
