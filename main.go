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
	ID        int
	Title     string
	Cards     []Card
	CreatedAt time.Time
	UpdatedAt time.Time
}
type Board struct {
	ID        int
	Title     string
	Lists     []List
	CreatedAt time.Time
	UpdatedAt time.Time
}

func remove(board *Board, listID int) {
	found := false
	for i, list := range board.Lists {
		if list.ID == listID {
			board.Lists = append(board.Lists[:i], board.Lists[i+1:]...)
			found = true
			fmt.Println("Лист успешно удален✅", list.Title)
			break
		}
	}
	if !found {
		fmt.Println("Список не найден❌")
	}
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
		if write == 1 {
			var title string
			fmt.Print("Введите название доски:")
			fmt.Scan(&title)
			newBoard := Board{
				ID:        boardID,
				Title:     title,
				Lists:     []List{},
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			boardID++
			board = append(board, newBoard)
			fmt.Println("Таблица создана")
		}
		if write == 2 {
			if len(board) == 0 {
				fmt.Println("Досок пока нету")
			}
			for _, board := range board {
				fmt.Println(board.Title)
			}
		}
		if write == 3 {
			if len(board) == 0 {
				fmt.Println("Досок пока нету")
				continue
			}
			for _, b := range board {
				fmt.Println(b.ID, b.Title)
			}
			fmt.Print("Введите номер доски:")
			var IDcheck int
			fmt.Scan(&IDcheck)
			var selectboard *Board
			for i := range board {
				if board[i].ID == IDcheck {
					selectboard = &board[i]
					break
				}
			}
			if selectboard == nil {
				fmt.Println("Доска с таким ID не найдена")
				continue
			}
			var IDlist int
			for {
				fmt.Println("Работа с доской:", selectboard.Title)
				fmt.Println("1. Добавить список")
				fmt.Println("2. Посмотреть списки")
				fmt.Println("3. Удалить список")
				fmt.Println("4. Управлять списком")
				fmt.Println("5. Вернуться в главное меню")
				fmt.Print("Выберите действие: ")
				fmt.Scan(&write)
				if write == 1 {
					var title string
					fmt.Print("Введите название списка")
					fmt.Scan(&title)
					newList := List{
						ID:    IDlist,
						Title: title,
						Cards: []Card{},
					}
					IDlist++
					selectboard.Lists = append(selectboard.Lists, newList)
					fmt.Println("Лист создан")
				}
				if write == 2 {
					if len(selectboard.Lists) == 0 {
						fmt.Println("Листов пока нету")
						continue
					}
					for _, l := range selectboard.Lists {
						fmt.Println(l.Title)
					}
				}
				if write == 3 {
					for _, l := range selectboard.Lists {
						fmt.Println(l.Title, l.ID)
					}
					fmt.Println("Введите ID листа, который вы хотите удалить")
					var DeleteID int
					fmt.Scan(&DeleteID)
					remove(selectboard, DeleteID)
				}
				if write == 5 {
					fmt.Println("Переход в главное меню")
					break
				}
			}
		}
		if write == 0 {
			return
		}
	}
}
