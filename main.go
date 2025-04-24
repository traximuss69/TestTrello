package main

import (
	"awesomeProject2/model"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var board []model.Board
	var list []model.List
	var cards []model.Card
	boardID := 1
	listID := 1
	cardID := 1
	http.HandleFunc("/boards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			for _, b := range board {
				fmt.Fprintln(w, b.Title)
			}
		} else if r.Method == http.MethodPost {
			var newBoard model.Board
			err := json.NewDecoder(r.Body).Decode(&newBoard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newBoard.ID = boardID
			newBoard.CreatedAt = time.Now()
			newBoard.UpdatedAt = time.Now()
			boardID++
			board = append(board, newBoard)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newBoard)
		}
	})
	http.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var newList model.List
			err := json.NewDecoder(r.Body).Decode(&newList)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newList.CreatedAt = time.Now()
			newList.UpdatedAt = time.Now()
			listID++
			for _, b := range board {
				if b.ID == newList.BoardID {
					b.Lists = append(b.Lists, newList)
				}
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(list)
		}
	})
	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var newCard model.Card
			var newList model.List
			err := json.NewDecoder(r.Body).Decode(&newCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newCard.ID = cardID
			newCard.CreatedAt = time.Now()
			newCard.UpdatedAt = time.Now()
			cardID++
		zalopka:
			for _, b := range board {
				if b.ID == newList.BoardID {
					for l := range b.Lists {
						if b.Lists[l].ID == newCard.ListID {
							b.Lists[l].Cards = append(b.Lists[l].Cards, newCard)
							break zalopka
						}
					}
				}
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(cards)
		}
		if r.Method == http.MethodPut {
			var updatedCard model.Card
			if err := json.NewDecoder(r.Body).Decode(&updatedCard); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			found := false
		zaloop:
			for _, b := range board {
				if b.ID == updatedCard.BoardID {
					for l := range b.Lists {
						if b.Lists[l].ID == updatedCard.ID {
							for j := range b.Lists[l].Cards {
								if b.Lists[l].Cards[j].ID == updatedCard.ID {
									b.Lists[l].Cards[j].Title = updatedCard.Title
									b.Lists[l].Cards[j].Description = updatedCard.Description
									b.Lists[l].Cards[j].UpdatedAt = time.Now()
									found = true
									json.NewEncoder(w).Encode(b.Lists[l].Cards[j])
									break zaloop
								}
							}
						}

					}
				}
			}
			if !found {
				http.Error(w, "Карточка не найдена", http.StatusNotFound)
			}
		}
		if r.Method == http.MethodDelete {
			var deletedCard model.Card
			if err := json.NewDecoder(r.Body).Decode(&deletedCard); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			found := false
		loop:
			for b := range board {
				if board[b].ID == deletedCard.BoardID {
					for l := range board[b].Lists {
						if board[b].Lists[l].ID == deletedCard.ID {
							for c := range board[b].Lists[l].Cards {
								if cards[c].ID == deletedCard.ID {
									board[b].Lists[l].RemoveCard(cardID)
									found = true
									w.WriteHeader(http.StatusOK)
									break loop
								}
							}
						}
					}
				}

			}
			if !found {
				http.Error(w, "Карточка не найдена", http.StatusNotFound)
			} else {
				fmt.Fprintln(w, "Карточка удалена", http.StatusOK)
			}

		}

	})
	http.ListenAndServe(":8080", nil)
}
