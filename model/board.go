package model

import "time"

type Board struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Lists      []List    `json:"lists"`
	NextListID int       `json:"next_list_id"`
	NextCardID int       `json:"next_card_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
type BoardDTO struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func BoardToDTO(b Board) BoardDTO {
	return BoardDTO{
		ID:    b.ID,
		Title: b.Title,
	}
}
