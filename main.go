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
	boardID := 1
	listID := 1
	cardID := 1
	http.HandleFunc("/boards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			for _, b := range board {
				fmt.Fprintln(w, b.Title)
			}
		}
		if r.Method == http.MethodPost {
			var newBoard model.Board
			err := json.NewDecoder(r.Body).Decode(&newBoard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newBoard.ID = boardID
			boardID++
			newBoard.CreatedAt = time.Now()
			newBoard.UpdatedAt = time.Now()
			board = append(board, newBoard)
			w.WriteHeader(http.StatusCreated)
			model.SendFullJSON(w, board)
		}
	})
	http.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			for _, b := range board {
				for _, l := range b.Lists {
					fmt.Fprintln(w, l.Title)
				}
			}
		}
		if r.Method == http.MethodPost {
			var newList model.List
			err := json.NewDecoder(r.Body).Decode(&newList)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newList.ID = listID
			listID++
			newList.CreatedAt = time.Now()
			newList.UpdatedAt = time.Now()
			for i := range board {
				if board[i].ID == newList.BoardID {
					board[i].Lists = append(board[i].Lists, newList)
				}
			}
			w.WriteHeader(http.StatusCreated)
			model.SendFullJSON(w, board)
		}
	})
	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			for _, b := range board {
				for _, l := range b.Lists {
					for _, c := range l.Cards {
						fmt.Fprintln(w, c.Title)
					}
				}
			}
		}
		if r.Method == http.MethodPost {
			var newCard model.Card
			err := json.NewDecoder(r.Body).Decode(&newCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newCard.ID = cardID
			cardID++
			newCard.CreatedAt = time.Now()
			newCard.UpdatedAt = time.Now()
		zalopka:
			for i := range board {
				if board[i].ID == newCard.BoardID {
					for l := range board[i].Lists {
						if board[i].Lists[l].ID == newCard.ListID {
							board[i].Lists[l].Cards = append(board[i].Lists[l].Cards, newCard)
							break zalopka
						}
					}
				}
			}
			w.WriteHeader(http.StatusCreated)
			model.SendFullJSON(w, board)
		} else if r.Method == http.MethodPut {
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
						if b.Lists[l].ID == updatedCard.ListID {
							for j := range b.Lists[l].Cards {
								if b.Lists[l].Cards[j].ID == updatedCard.ID {
									b.Lists[l].Cards[j].Title = updatedCard.Title
									b.Lists[l].Cards[j].Description = updatedCard.Description
									b.Lists[l].Cards[j].UpdatedAt = time.Now()
									found = true
									json.NewEncoder(w).Encode(b.Lists[l].Cards[j])
									model.SendFullJSON(w, board)
									break zaloop
								}
							}
						}
					}
				}
				if !found {
					http.Error(w, "Карточка не найдена", http.StatusNotFound)
				}
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
						if board[b].Lists[l].ID == deletedCard.ListID {
							for c := range board[b].Lists[l].Cards {
								if board[b].Lists[l].Cards[c].ID == deletedCard.ID {
									board[b].Lists[l].RemoveCard(deletedCard.ID)
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
				model.SendFullJSON(w, board)
			}

		}

	})
	http.ListenAndServe(":8080", nil)
}
