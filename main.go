package main

import (
	"encoding/json"
	"fmt"
	"os"
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
func removeCard(list *List, cardID int) {
	found := false
	for i, card := range list.Cards {
		if card.ID == cardID {
			list.Cards = append(list.Cards[:i], list.Cards[i+1:]...)
			found = true
			fmt.Println("Карточка успешно удалена", card)
			break
		}
	}
	if !found {
		fmt.Println("Карточка не найдена")
	}
}
func moveCard(list *List, toList *List, cardID int) {
	var moveCard Card
	found := false
	for i, card := range list.Cards {
		if card.ID == cardID {
			moveCard = card
			list.Cards = append(list.Cards[:i], list.Cards[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Карточка не найдена")
		return
	}
	moveCard.Status = list.Title
	moveCard.UpdatedAt = time.Now()

	toList.Cards = append(toList.Cards, moveCard)
	fmt.Println("Вы успешно переместили карточку")
}
func saveToFile(boards []Board, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(boards)
}
func loadFromFile(filename string) ([]Board, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var boards []Board
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&boards)
	return boards, err
}
func editCard(card *Card) {
	var newTitle, newDescription, newStatus string
	fmt.Print("Введите новое название(оставьте пустым если не хотите менять):")
	fmt.Scan(&newTitle)
	if newTitle != "" {
		card.Title = newTitle
	}
	fmt.Print("Введите новое описание(оставьте пустым если не хотите менять):")
	fmt.Scan(&newDescription)
	if newDescription != "" {
		card.Description = newDescription
	}
	fmt.Print("Введите новый статус(оставьте пустым если не хотите менять):")
	fmt.Scan(&newStatus)
	if newStatus != "" {
		card.Status = newStatus
	}
	card.UpdatedAt = time.Now()
	fmt.Println("Карточка успешно обновлена✅")
}
func main() {
	var board []Board
	var boardID int
	var write int
	var cardID int
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
			fmt.Println("Таблица создана✅")
		}
		if write == 2 {
			if len(board) == 0 {
				fmt.Println("Досок пока нету❌")
			}
			for _, board := range board {
				fmt.Println(board.Title)
			}
		}
		if write == 3 {
			if len(board) == 0 {
				fmt.Println("Досок пока нету❌")
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
				fmt.Println("Доска с таким ID не найдена❌")
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
					fmt.Println("Лист создан✅")
				}
				if write == 2 {
					if len(selectboard.Lists) == 0 {
						fmt.Println("Листов пока нету❌")
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
				if write == 4 {
					if len(selectboard.Lists) == 0 {
						fmt.Println("Список отсуствует❌")
						continue
					}
					for _, l := range selectboard.Lists {
						fmt.Println(l.ID, l.Title)
					}
					fmt.Println("Введите ID списка")
					var listcheck int
					fmt.Scan(&listcheck)

					var selectlist *List
					for i := range selectboard.Lists {
						if selectboard.Lists[i].ID == listcheck {
							selectlist = &selectboard.Lists[i]
							break
						}
					}
					if selectlist == nil {
						fmt.Println("Вы ввели неправильный ID❌")
						break
					}

					for {
						fmt.Println("📂 Меню управления списком")
						fmt.Println("Вы выбрали список: \"In Progress\"")
						fmt.Println("Что вы хотите сделать?")
						fmt.Println("1. Посмотреть карточки")
						fmt.Println("2. Добавить карточку")
						fmt.Println("3. Удалить карточку")
						fmt.Println("4. Переместить карточку в другой список")
						fmt.Println("5. Редактировать карточку")
						fmt.Println("6. Вернуться к доске")
						fmt.Println("Выберите действие:")
						fmt.Scan(&write)
						if write == 1 {
							if len(selectlist.Cards) == 0 {
								fmt.Println("Карточки отсуствуют❌")
							} else {
								fmt.Println("Карточки:")
								for _, card := range selectlist.Cards {
									fmt.Println(card.ID, card.Title)
								}
							}
						}
						if write == 2 {
							var title string
							fmt.Print("Введите название карточки")
							fmt.Scan(&title)
							newCard := Card{
								ID:        cardID,
								Title:     title,
								Status:    "In Progress",
								CreatedAt: time.Now(),
								UpdatedAt: time.Now(),
							}
							cardID++
							selectlist.Cards = append(selectlist.Cards, newCard)
							fmt.Println("Карточка создана✅")
						}
						if write == 3 {
							var deleteCard int
							if len(selectlist.Cards) == 0 {
								fmt.Println("Карта отсуствует❌")
							} else {
								for _, card := range selectlist.Cards {
									fmt.Println(card.ID, card.Title)
								}
								fmt.Println("Введите ID карточки, который вы хотите удалить")
								fmt.Scan(&deleteCard)
								removeCard(selectlist, deleteCard)
							}
						}
						if write == 4 {
							var cardID int
							var selectIDlist int
							for _, card := range selectlist.Cards {
								fmt.Println(card.ID, card.Title)
							}
							fmt.Print("Введите ID карточки:")
							fmt.Scan(&cardID)

							fmt.Print("Введите ID списка:")
							fmt.Scan(&selectIDlist)

							var toList *List
							for i := range selectlist.Cards {
								if selectboard.Lists[i].ID == selectIDlist {
									toList = &selectboard.Lists[i]
									break
								}
							}
							if selectlist == nil {
								fmt.Println("Неправильный ввод❌")
							} else {
								moveCard(selectlist, toList, cardID)
							}
						}
						if write == 5 {
							if len(selectlist.Cards) == 0 {
								fmt.Println("Карта отсуствует❌")
								continue
							}
							for _, card := range selectlist.Cards {
								fmt.Println(card.ID, card.Title)
							}
							fmt.Print("Введи ID карточки")
							var cardID int
							fmt.Scan(&cardID)
							var selectCard *Card
							for i := range selectlist.Cards {
								if selectlist.Cards[i].ID == cardID {
									selectCard = &selectlist.Cards[i]
									break
								}
							}
							if selectCard == nil {
								fmt.Println("Карточка не найдена❌")
							} else {
								editCard(selectCard)
							}
						}
						if write == 6 {
							fmt.Println("Переход к доске🔙")
							break
						}
					}

				}
				if write == 5 {
					fmt.Println("Переход в главное меню🔙")
					break
				}
			}
		}
		if write == 4 {
			var filename string
			fmt.Print("Введите имя файла для загрузки: ")
			fmt.Scan(&filename)
			loadedBoards, err := loadFromFile(filename)
			if err != nil {
				fmt.Println("Ошибка при загрузке:", err)
			} else {
				board = loadedBoards
				fmt.Println("✅ Данные успешно загружены!")
			}
		}

		if write == 5 {
			var filename string
			fmt.Print("Введите имя файла для сохранения: ")
			fmt.Scan(&filename)
			err := saveToFile(board, filename)
			if err != nil {
				fmt.Println("Ошибка при сохранении:", err)
			} else {
				fmt.Println("💾 Данные успешно сохранены!")
			}
		}

		if write == 0 {
			fmt.Println("Ещё увидиммся")
			return
		}
	}
}
