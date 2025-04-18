package model

import (
	"fmt"
	"time"
)

type Card struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
type List struct {
	ID        int
	Title     string
	Cards     []Card
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Board struct {
	ID         int
	Title      string
	Lists      []List
	NextListID int
	NextCardID int
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
