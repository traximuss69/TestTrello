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
