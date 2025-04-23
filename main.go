package main

import (
	"awesomeProject2/model"
	//"awesomeProject2/storage"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var board []model.Board
	var list []model.List
	var cards []model.Card
	boardID := 1
	listID := 1
	cardID := 1
	http.HandleFunc("/boards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			for _, b := range board {
				fmt.Fprintln(w, b.Title)
			}
		} else if r.Method == http.MethodPost {
			var newBoard model.Board
			err := json.NewDecoder(r.Body).Decode(&newBoard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newBoard.ID = boardID
			newBoard.CreatedAt = time.Now()
			newBoard.UpdatedAt = time.Now()
			boardID++
			board = append(board, newBoard)
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newBoard)
		}
	})
	http.HandleFunc("/lists", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var newList model.List
			err := json.NewDecoder(r.Body).Decode(&newList)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newList.CreatedAt = time.Now()
			newList.UpdatedAt = time.Now()
			listID++
			for _, b := range board {
				if b.ID == newList.BoardID {
					b.Lists = append(b.Lists, newList)
				}
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(list)
		}
	})
	http.HandleFunc("/cards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var newCard model.Card
			err := json.NewDecoder(r.Body).Decode(&newCard)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			newCard.ID = cardID
			newCard.CreatedAt = time.Now()
			newCard.UpdatedAt = time.Now()
			cardID++
			for _, l := range list {
				if l.ID == newCard.ListID {
					l.Cards = append(l.Cards, newCard)
				}
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(cards)
		}
		if r.Method == http.MethodPut {
			var updatedCard model.Card
			if err := json.NewDecoder(r.Body).Decode(&updatedCard); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			found := false
			for i := range cards {
				if cards[i].ID == updatedCard.ID {
					cards[i].Title = updatedCard.Title
					cards[i].Description = updatedCard.Description
					cards[i].UpdatedAt = time.Now()
					found = true
					json.NewEncoder(w).Encode(cards[i])
					break
				}
			}
			if !found {
				http.Error(w, "Карточка не найдена", http.StatusNotFound)
			}
		}
		if r.Method == http.MethodDelete {
			var deletedCard model.Card
			found := false
			for i := range cards {
				if cards[i].ID == deletedCard.ID {
					cards = append(cards[:i], cards[i+1:]...)
					fmt.Fprintln(w, "Карточка удалена")
					found = true
					break
				}
				if !found {
					http.Error(w, "Карточка не найдена", http.StatusNotFound)
				}
			}
		}
	})
	http.ListenAndServe(":8080", nil)
}

//	for {
//		fmt.Println("Выберите действие:")
//		fmt.Println("1. Создать Доску")
//		fmt.Println("2. Посмотреть все доски")
//		fmt.Println("3. Управлять доской")
//		fmt.Println("4. Загрузить данные из файла")
//		fmt.Println("5. Сохранить данные в файл")
//		fmt.Println("0. Выход")
//		fmt.Print("Введите число:")
//		fmt.Scan(&write)
//		if write == 1 {
//			var title string
//			fmt.Print("Введите название доски:")
//			fmt.Scan(&title)
//			newBoard := model.Board{
//				ID:        boardID,
//				Title:     title,
//				Lists:     []model.List{},
//				CreatedAt: time.Now(),
//				UpdatedAt: time.Now(),
//			}
//			boardID++
//			board = append(board, newBoard)
//			fmt.Println("Таблица создана✅")
//		}
//		if write == 2 {
//			if len(board) == 0 {
//				fmt.Println("Досок пока нету❌")
//			}
//			for _, board := range board {
//				fmt.Println(board.Title)
//			}
//		}
//		if write == 3 {
//			if len(board) == 0 {
//				fmt.Println("Досок пока нету❌")
//				continue
//			}
//			for _, b := range board {
//				fmt.Println(b.ID, b.Title)
//			}
//			fmt.Print("Введите номер доски:")
//			var IDcheck int
//			fmt.Scan(&IDcheck)
//			var selectboard *model.Board
//			for i := range board {
//				if board[i].ID == IDcheck {
//					selectboard = &board[i]
//					break
//				}
//			}
//			if selectboard == nil {
//				fmt.Println("Доска с таким ID не найдена❌")
//				continue
//			}
//			for {
//				fmt.Println("Работа с доской:", selectboard.Title)
//				fmt.Println("1. Добавить список")
//				fmt.Println("2. Посмотреть списки")
//				fmt.Println("3. Удалить список")
//				fmt.Println("4. Управлять списком")
//				fmt.Println("5. Вернуться в главное меню")
//				fmt.Print("Выберите действие: ")
//				var writeBoard int
//				fmt.Scan(&writeBoard)
//				if writeBoard == 1 {
//					var title string
//					fmt.Print("Введите название списка: ")
//					fmt.Scan(&title)
//					newList := model.List{
//						ID:    selectboard.NextListID,
//						Title: title,
//						Cards: []model.Card{},
//					}
//					selectboard.NextListID++
//					selectboard.Lists = append(selectboard.Lists, newList)
//					fmt.Println("Лист создан✅")
//				}
//				if writeBoard == 2 {
//					if len(selectboard.Lists) == 0 {
//						fmt.Println("Листов пока нету❌")
//						continue
//					}
//					for _, l := range selectboard.Lists {
//						fmt.Println(l.Title)
//					}
//				}
//				if writeBoard == 3 {
//					for _, l := range selectboard.Lists {
//						fmt.Println(l.Title, l.ID)
//					}
//					fmt.Print("Введите ID листа, который вы хотите удалить")
//					var DeleteID int
//					fmt.Scan(&DeleteID)
//					selectboard.RemoveList(DeleteID)
//				}
//				if writeBoard == 4 {
//					if len(selectboard.Lists) == 0 {
//						fmt.Println("Список отсуствует❌")
//						continue
//					}
//					for _, l := range selectboard.Lists {
//						fmt.Println(l.ID, l.Title)
//					}
//					fmt.Print("Введите ID списка")
//					var listcheck int
//					fmt.Scan(&listcheck)
//
//					var selectlist *model.List
//					for i := range selectboard.Lists {
//						if selectboard.Lists[i].ID == listcheck {
//							selectlist = &selectboard.Lists[i]
//							break
//						}
//					}
//					if selectlist == nil {
//						fmt.Println("Вы ввели неправильный ID❌")
//						break
//					}
//					for {
//						fmt.Println("📂 Меню управления списком")
//						fmt.Println("Вы выбрали список:", selectlist.Title)
//						fmt.Println("Что вы хотите сделать?")
//						fmt.Println("1. Посмотреть карточки")
//						fmt.Println("2. Добавить карточку")
//						fmt.Println("3. Удалить карточку")
//						fmt.Println("4. Переместить карточку в другой список")
//						fmt.Println("5. Редактировать карточку")
//						fmt.Println("6. Вернуться к доске")
//						fmt.Print("Выберите действие:")
//						var writeCart int
//						fmt.Scan(&writeCart)
//						if writeCart == 1 {
//							if len(selectlist.Cards) == 0 {
//								fmt.Println("Карточки отсуствуют❌")
//							} else {
//								fmt.Println("Карточки: ")
//								for _, card := range selectlist.Cards {
//									fmt.Println(card.ID, card.Title)
//								}
//							}
//						}
//						if writeCart == 2 {
//							var title string
//							fmt.Print("Введите название карточки")
//							fmt.Scan(&title)
//							newCard := model.Card{
//								ID:        selectboard.NextCardID,
//								Title:     title,
//								Status:    selectlist.Title,
//								CreatedAt: time.Now(),
//								UpdatedAt: time.Now(),
//							}
//							selectboard.NextCardID++
//							selectlist.Cards = append(selectlist.Cards, newCard)
//							fmt.Println("Карточка создана✅")
//						}
//						if writeCart == 3 {
//							var deleteCard int
//							if len(selectlist.Cards) == 0 {
//								fmt.Println("Карта отсуствует❌")
//							} else {
//								for _, card := range selectlist.Cards {
//									fmt.Println(card.ID, card.Title)
//								}
//								fmt.Print("Введите ID карточки, который вы хотите удалить")
//								fmt.Scan(&deleteCard)
//								selectlist.RemoveCard(deleteCard)
//							}
//						}
//						if writeCart == 4 {
//							if len(selectboard.Lists) < 2 {
//								fmt.Println("Должно быть хотя бы 2 списка")
//								continue
//							}
//							if len(selectlist.Cards) <= 0 {
//								fmt.Println("Карточки отсуствуют")
//								continue
//							}
//							var cardID int
//							var selectIDlist int
//							fmt.Print("Карточки в списке: ")
//							for _, l := range selectlist.Cards {
//								fmt.Println(l.ID, l.Title)
//							}
//							fmt.Print("Введите ID карточки: ")
//							fmt.Scan(&cardID)
//							fmt.Println("Весь список: ")
//							for _, l := range selectboard.Lists {
//								fmt.Println(l.ID, l.Title)
//							}
//							fmt.Print("Введите ID списка, в который хотите переместить: ")
//							fmt.Scan(&selectIDlist)
//
//							var toList *model.List
//							for i := range selectboard.Lists {
//								if selectboard.Lists[i].ID == selectIDlist {
//									toList = &selectboard.Lists[i]
//									break
//								}
//							}
//							if toList == nil {
//								fmt.Println("Неправильный ввод❌")
//								continue
//							}
//							if toList.ID == selectlist.ID {
//								fmt.Println("Нельзя переместить в тот же список❌")
//								continue
//							}
//							selectlist.MoveCard(toList, cardID)
//						}
//						if writeCart == 5 {
//							if len(selectlist.Cards) == 0 {
//
//								fmt.Println("Карта отсуствует❌")
//								continue
//							}
//							for _, card := range selectlist.Cards {
//								fmt.Println(card.ID, card.Title)
//							}
//							fmt.Print("Введи ID карточки")
//							var cardID int
//							fmt.Scan(&cardID)
//							var selectCard *model.Card
//							for i := range selectlist.Cards {
//								if selectlist.Cards[i].ID == cardID {
//									selectCard = &selectlist.Cards[i]
//									break
//								}
//							}
//							if selectCard == nil {
//								fmt.Println("Карточка не найдена❌")
//							} else {
//								selectCard.Edit()
//							}
//						}
//						if writeCart == 6 {
//							fmt.Println("Переход к доске🔙")
//							break
//						}
//					}
//				}
//				if writeBoard == 5 {
//					fmt.Println("Переход в главное меню🔙")
//					break
//				}
//			}
//		}
//		if write == 4 {
//			var filename string
//			fmt.Print("Введите имя файла для загрузки: ")
//			fmt.Scan(&filename)
//			loadedBoards, err := storage.LoadFromFile(filename)
//			if err != nil {
//				fmt.Println("Ошибка при загрузке:", err)
//			} else {
//				board = loadedBoards
//				fmt.Println("✅ Данные успешно загружены!")
//			}
//		}
//
//		if write == 5 {
//			var filename string
//			fmt.Print("Введите имя файла для сохранения: ")
//			fmt.Scan(&filename)
//			err := storage.SaveToFile(board, filename)
//			if err != nil {
//				fmt.Println("Ошибка при сохранении:", err)
//			} else {
//				fmt.Println("💾 Данные успешно сохранены!")
//			}
//		}
//
//		if write == 0 {
//			fmt.Println("Ещё увидиммся")
//			return
//		}
//	}
//}
