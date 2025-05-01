package handler

import (
	"awesomeProject2/service"
	"encoding/json"
	"net/http"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{Service: service}
}
func (h *Handler) RegisterRouters() {
	http.HandleFunc("/boards", h.handleBoards)
	http.HandleFunc("/lists", h.handleLists)
	http.HandleFunc("/cards", h.handleCards)
}
func (h *Handler) handleBoards(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		titles := h.Service.GetBoards()
		json.NewEncoder(w).Encode(titles)
	} else if r.Method == http.MethodPost {
		var input struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		board := h.Service.CreateBoard(input.Title)
		json.NewEncoder(w).Encode(board)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
