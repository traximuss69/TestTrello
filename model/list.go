package model

import "time"

type List struct {
	ID        int       `json:"id"`
	BoardID   int       `json:"board_id"`
	Title     string    `json:"title"`
	Cards     []Card    `json:"cards"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type ListDTO struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	BoardID int    `json:"board_id"`
}

func ListToDTO(l List) ListDTO {
	return ListDTO{
		ID:      l.ID,
		Title:   l.Title,
		BoardID: l.BoardID,
	}
}
