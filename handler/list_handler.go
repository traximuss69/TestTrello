package handler

import (
	"awesomeProject2/cmd/model"
	"awesomeProject2/cmd/service"
	"encoding/json"
	"net/http"
)

type ListHandler struct {
	Service *service.Service
}

func NewListHandler(service *service.Service) *ListHandler {
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
		lists := h.Service.GetList(listID)
		var dto []model.ListDTO
		for i := range lists {
			dto = append(dto, model.ListToDTO(lists[i]))
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dto); err != nil {
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
		list := h.Service.CreateList(input.BoardID, input.Title)
		dto := model.ListToDTO(list)
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
