package storage

import (
	"awesomeProject2/model"
	"encoding/json"
	"os"
)

func SaveToFile(boards []model.Board, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(boards)
}

func LoadFromFile(filename string) ([]model.Board, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var boards []model.Board
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&boards)
	return boards, err
}
