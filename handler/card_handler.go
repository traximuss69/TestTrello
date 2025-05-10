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
		var requestBody dto.CardDTO
		if r.Body != nil {
			defer r.Body.Close()
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
		}
		var cardID *int
		if requestBody.ID != 0 {
			cardID = &requestBody.ID
		}
		cards := h.service.GetCards(cardID)
		var cardDTOs []dto.CardDTO
		for i := range cards {
			cardDTOs = append(cardDTOs, dto.CardToDTO(cards[i]))
		}
		if err := json.NewEncoder(w).Encode(cardDTOs); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	} else if r.Method == http.MethodPost {
		var input dto.CreateCardDTO
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if input.BoardID == 0 || input.ListID == 0 {
			http.Error(w, "board id or list id required", http.StatusBadRequest)
		}
		if input.Title == "" {
			http.Error(w, "title is required", http.StatusBadRequest)
		}
		card := h.service.CreateCard(input.Title, input.BoardID, input.ListID, input.Description)
		cardDTOs := dto.CardToDTO(card)
		json.NewEncoder(w).Encode(cardDTOs)
	} else if r.Method == http.MethodDelete {
		var input dto.DeleteCardDTO
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		if input.CardID == 0 || input.ListID == 0 || input.BoardID == 0 {
			http.Error(w, "card id or list id required", http.StatusBadRequest)
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
		var updatedCardDTO dto.CardDTO
		if err := json.NewDecoder(r.Body).Decode(&updatedCardDTO); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updatedCard := model.Card{
			Title:       updatedCardDTO.Title,
			Description: updatedCardDTO.Description,
		}
		if updatedCardDTO.BoardID == 0 || updatedCardDTO.ListID == 0 || updatedCardDTO.ID == 0 {
			http.Error(w, "id error", http.StatusBadRequest)
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
