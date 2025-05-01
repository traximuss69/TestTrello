package main

import (
	"awesomeProject2/handler"
	"awesomeProject2/service"
	"awesomeProject2/storage"
	"net/http"
)

func main() {
	store := storage.NewStorage()
	srv := service.NewService(store)
	h := handler.NewHandler(srv)
	h.RegisterRouters()
	http.ListenAndServe(":8080", nil)
}
