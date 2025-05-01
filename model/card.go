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
