package model

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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
type List struct {
	ID        int       `json:"id"`
	BoardID   int       `json:"board_id"`
	Title     string    `json:"title"`
	Cards     []Card    `json:"cards"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Board struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Lists      []List    `json:"lists"`
	NextListID int       `json:"next_list_id"`
	NextCardID int       `json:"next_card_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (b *Board) RemoveList(listID int) {
	found := false
	for i, list := range b.Lists {
		if list.ID == listID {
			b.Lists = append(b.Lists[:i], b.Lists[i+1:]...)
			found = true
			fmt.Println("Лист успешно удален✅", list.Title)
			break
		}
	}
	if !found {
		fmt.Println("Список не найден❌")
	}
}
func (l *List) RemoveCard(cardID int) (Card, bool) {
	for i, card := range l.Cards {
		if card.ID == cardID {
			l.Cards = append(l.Cards[:i], l.Cards[i+1:]...)
			fmt.Println("Карточка успешно удалена", card)
			return card, true
		}
	}
	fmt.Println("Карточка не найдена")
	return Card{}, false
}
func (l *List) MoveCard(toList *List, cardID int) {
	moveCard, found := l.RemoveCard(cardID)
	if !found {
		fmt.Println("Карточка не найдена❌")
		return
	}
	toList.Cards = append(toList.Cards, moveCard)
	fmt.Println("Карточка успешно перемещена✅", moveCard.Title)
}
func (card *Card) Edit() {
	var newTitle, newDescription string
	fmt.Print("Введите новое название(оставьте пустым если не хотите менять):")
	fmt.Scan(&newTitle)
	if newTitle != "" {
		card.Title = newTitle
	}
	fmt.Print("Введите новое описание(оставьте пустым если не хотите менять):")
	fmt.Scan(&newDescription)
	if newDescription != "" {
		card.Description = newDescription
	}
	card.UpdatedAt = time.Now()
	fmt.Println("Карточка успешно обновлена✅")
}
func SendFullJSON(w http.ResponseWriter, board []Board) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(board)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
