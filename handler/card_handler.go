package handler

import (
	"awesomeProject2/cmd/model"
	"awesomeProject2/cmd/service"
	"encoding/json"
	"net/http"
)

type CardHandler struct {
	Service *service.Service
}

func NewCardHandler(service *service.Service) *CardHandler {
	return &CardHandler{service}
}
func (h *CardHandler) HandleCards(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var cardID *int
		if r.Body != nil {
			var requestBody struct {
				ID int `json:"id"`
			}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if requestBody.ID != 0 {
				cardID = &requestBody.ID
			}
		}
		cards := h.Service.GetCard(cardID)
		var dto []model.CardDTO
		for i := range cards {
			dto = append(dto, model.CardToDTO(cards[i]))
		}
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
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
		dto := model.CardToDTO(card)
		json.NewEncoder(w).Encode(dto)
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
		deletedCard, ok := h.Service.DeleteCard(input.BoardID, input.ListID, input.CardID)
		if !ok {
			http.Error(w, "Error deleting card", http.StatusInternalServerError)
			return
		}
		dto := model.CardToDTO(deletedCard)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(dto); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPut {
		var updatedCardDTO model.UpdatedCardDTO
		if err := json.NewDecoder(r.Body).Decode(&updatedCardDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedCard := model.Card{
			Title:       updatedCardDTO.Title,
			Description: updatedCardDTO.Description,
		}
		updatedCard, ok := h.Service.UpdatedCard(updatedCard)
		if !ok {
			http.Error(w, "Error updating card", http.StatusInternalServerError)
			return
		}
		updatedCardDTOResponse := model.CardToDTO(updatedCard)
		json.NewEncoder(w).Encode(updatedCardDTOResponse)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
