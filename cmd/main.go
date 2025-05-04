package main

import (
	"awesomeProject2/cmd/handler"
	"awesomeProject2/cmd/service"
	"awesomeProject2/cmd/storage"
	"net/http"
)

func main() {
	store := storage.NewStorage()
	srv := service.NewService(store)
	boardHandler := handler.NewBoardHandler(srv)
	listHandler := handler.NewListHandler(srv)
	cardHandler := handler.NewCardHandler(srv)
	http.HandleFunc("/boards", boardHandler.HandleBoards)
	http.HandleFunc("/lists", listHandler.HandleLists)
	http.HandleFunc("/cards", cardHandler.HandleCards)
	http.ListenAndServe(":8080", nil)
}
