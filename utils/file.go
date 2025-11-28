package utils

import (
	"encoding/json"
	"errors"
	"os"
	"mini_project2/model"
)

const TodoFilePath = "data/todo-lists.json"

func EnsureTodosFile() error {
	_, err := os.Stat(TodoFilePath)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll("data", 0755); err != nil {
			return err
		}
		return os.WriteFile(TodoFilePath, []byte("[]"), 0644)
	}
	return nil
}

func ReadTodosFromFile() ([]model.Todo, error) {
	if err := EnsureTodosFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(TodoFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo

	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, err
	}
	return todos, nil
}

func WriteTodosToFile(todos []model.Todo) error {
	bytes, err := json.MarshalIndent(todos, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(TodoFilePath, bytes, 0644)
}