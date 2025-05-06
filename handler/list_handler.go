package handler

import (
	"awesomeProject2/cmd/dto"
	"encoding/json"
	"net/http"
)

type ListHandler struct {
	service ListService
}

func NewListHandler(service ListService) *ListHandler {
	return &ListHandler{service}
}
func (h *ListHandler) HandleLists(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var listID *int
		if r.Body != nil {
			var requestBody struct {
				ID int `json:"id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if requestBody.ID != 0 {
				listID = &requestBody.ID
			}
		}
		lists := h.service.GetLists(listID)
		var ListDTOs []dto.ListDTO
		for i := range lists {
			ListDTOs = append(ListDTOs, dto.ListToDTO(lists[i]))
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(ListDTOs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		var input struct {
			BoardID int    `json:"board_id"`
			Title   string `json:"title"`
		}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		list := h.service.CreateList(input.Title, input.BoardID)
		dto := dto.ListToDTO(list)
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
