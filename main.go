package main

import (
	"fmt"
	"time"
)

type Card struct {
	ID          int
	Title       string
	Description string
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type List struct {
	ID    int
	Title string
	Cards []Card
}
type Board struct {
	ID    int
	Title string
	Lists []List
}

func main() {
	var board []Board
	var boardID int
	var write int
	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1. Создать Доску")
		fmt.Println("2. Посмотреть все доски")
		fmt.Println("3. Управлять доской")
		fmt.Println("4. Загрузить данные из файла")
		fmt.Println("5. Сохранить данные в файл")
		fmt.Println("0. Выход")
		fmt.Print("Введите число:")
		fmt.Scan(&write)
	}
}
