package main

import (
	"awesomeProject2/cmd/config"
	"awesomeProject2/cmd/db"
	"awesomeProject2/cmd/handler"
	"awesomeProject2/cmd/service"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

func main() {
	env := config.GetOrDefault()

	db, err := sqlx.Open("postgres", env)
	if err != nil {
		log.Fatalf("не удалось подключиться к БД: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("пинг БД не прошёл: %v", err)
	}
	store := storage.NewStorage(db)
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
