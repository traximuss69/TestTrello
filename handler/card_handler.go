package handler

import (
	"awesomeProject2/cmd/dto"
	"awesomeProject2/cmd/model"
	"encoding/json"
	"net/http"
)

type CardHandler struct {
	service CardService
}

func NewCardHandler(service CardService) *CardHandler {
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
		cards := h.service.GetCards(cardID)
		var CardDTOs []dto.CardDTO
		for i := range cards {
			CardDTOs = append(CardDTOs, dto.CardToDTO(cards[i]))
		}
		if err := json.NewEncoder(w).Encode(CardDTOs); err != nil {
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
		card := h.service.CreateCard(input.Title, input.BoardID, input.ListID, input.Description)
		CardDTOs := dto.CardToDTO(card)
		json.NewEncoder(w).Encode(CardDTOs)
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
		deletedCard, err := h.service.DeleteCard(input.BoardID, input.ListID, input.CardID)
		if err != nil {
			http.Error(w, "Error deleting card", http.StatusInternalServerError)
			return
		}
		CardDTOs := dto.CardToDTO(deletedCard)
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(CardDTOs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPut {
		var updatedCardDTO dto.UpdatedCardDTO
		if err := json.NewDecoder(r.Body).Decode(&updatedCardDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedCard := model.Card{
			Title:       updatedCardDTO.Title,
			Description: updatedCardDTO.Description,
		}
		updatedCard, err := h.service.UpdateCard(updatedCard)
		if err != nil {
			http.Error(w, "Error updating card", http.StatusInternalServerError)
			return
		}
		updatedCardDTOResponse := dto.CardToDTO(updatedCard)
		json.NewEncoder(w).Encode(updatedCardDTOResponse)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
