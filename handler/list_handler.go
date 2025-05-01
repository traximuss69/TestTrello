package handler

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) handleLists(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		titles := h.Service.GetList()
		json.NewEncoder(w).Encode(titles)
		return
	} else if r.Method == http.MethodPost {
		var input struct {
			BoardID int    `json:"board_id"`
			Title   string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		list := h.Service.CreateList(input.BoardID, input.Title)
		json.NewEncoder(w).Encode(list)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
