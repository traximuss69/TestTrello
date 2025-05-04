package handler

import (
	"awesomeProject2/cmd/model"
	"awesomeProject2/cmd/service"
	"encoding/json"
	"net/http"
)

type BoardHandler struct {
	Service *service.Service
}

func NewBoardHandler(service *service.Service) *BoardHandler {
	return &BoardHandler{service}
}
func (h *BoardHandler) HandleBoards(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var boardID *int
		if r.Body != nil {
			var requestBody struct {
				ID int `json:"id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if requestBody.ID != 0 {
				boardID = &requestBody.ID
			}
		}
		boards := h.Service.GetBoards(boardID)
		var dto []model.BoardDTO
		for i := range boards {
			dto = append(dto, model.BoardToDTO(boards[i]))
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		var input struct {
			Title string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		board := h.Service.CreateBoard(input.Title)
		dto := model.BoardToDTO(board)
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
