package main

import (
	"awesomeProject2/cmd/handler"
	"awesomeProject2/cmd/service"
	"awesomeProject2/cmd/storage"
	"net/http"
)

func main() {
	store := storage.NewStorage()
	boardService := service.NewBoardService(store)
	listService := service.NewListService(store)
	cardService := service.NewCardService(store)
	boardHandler := handler.NewBoardHandler(boardService)
	listHandler := handler.NewListHandler(listService)
	cardHandler := handler.NewCardHandler(cardService)
	http.HandleFunc("/boards", boardHandler.HandleBoards)
	http.HandleFunc("/lists", listHandler.HandleLists)
	http.HandleFunc("/cards", cardHandler.HandleCards)
	http.ListenAndServe(":8080", nil)
}
