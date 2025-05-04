package model

import "time"

type Card struct {
	ID          int       `json:"id"`
	BoardID     int       `json:"board_id"`
	ListID      int       `json:"list_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
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

func CardToDTO(c Card) CardDTO {
	return CardDTO{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		ListID:      c.ListID,
	}
}
