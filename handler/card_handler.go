package handler

import (
	"awesomeProject2/model"
	"encoding/json"
	"net/http"
)

func (h *Handler) handleCards(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		titles := h.Service.GetCard()
		json.NewEncoder(w).Encode(titles)
		return
	} else if r.Method == http.MethodPost {
		var input struct {
			BoardID     int    `json:"board_id"`
			ListID      int    `json:"list_id"`
			Title       string `json:"title"`
			Description string `json:"description"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		card := h.Service.CreateCard(input.BoardID, input.ListID, input.Title, input.Description)
		json.NewEncoder(w).Encode(card)
	} else if r.Method == http.MethodDelete {
		var input struct {
			BoardID int `json:"board_id"`
			ListID  int `json:"list_id"`
			CardID  int `json:"card_id"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, ok := h.Service.DeleteCard(input.BoardID, input.ListID, input.CardID)
		if !ok {
			http.Error(w, "Error deleting card", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	} else if r.Method == http.MethodPut {
		var updated model.Card
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedCard, ok := h.Service.UpdatedCard(updated)
		if !ok {
			http.Error(w, "Error updating card", http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(updatedCard)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
