package service

import (
	"errors"
	"mini_project2/dto"
	"mini_project2/model"
	"mini_project2/utils"
	"strings"
	"time"
)

type TodoService struct {}

func NewTodoService() TodoService {
	return TodoService{}
}

func (t *TodoService) CreateTodo(req dto.CreateTodoRequest) (model.Todo, error) {
	if strings.TrimSpace(req.Task_name) == "" {
		return model.Todo{}, errors.New("task Name is Required")
	}

	if strings.TrimSpace(req.Priority) == "" {
		return model.Todo{}, errors.New("priority is required")
	}

	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return model.Todo{}, err
	}

	for _, t := range todos {
		if t.Task_name == req.Task_name {
			return model.Todo{}, errors.New("this task is already made")
		}
	}

	newID := getNextID(todos)

	for _, t := range todos {
		if int(t.Id) >= newID {
			newID = t.Id + 1
		}
 	}

	newTodos := model.Todo{
		Base: model.Base{
			Id: newID,
			CreatedAt: time.Now(),
		},
		Task_name: req.Task_name,
		Priority: req.Priority,
		Status: "pending",
	}

	todos = append(todos, newTodos)

	if err := utils.WriteTodosToFile(todos); err != nil {
		return model.Todo{}, err
	}

	return newTodos, nil
}

func (t *TodoService) ListTodo() (*[]dto.ListTodoResponse, error) {
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return nil, err
	}

	var listTodos []dto.ListTodoResponse

	for _, t := range todos {
		todo := dto.ListTodoResponse{
			Id: t.Id,
			Task_name: t.Task_name,
			Status: t.Status,
			Priority: t.Priority,
		}
		listTodos = append(listTodos, todo)
	}

	return &listTodos, err
}

func (t *TodoService) UpdateStatus(req dto.UpdateTodoRequest) (model.Todo, error) {
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return model.Todo{}, err
	}

	var updated model.Todo
	found := false

	for i, t := range todos {
		if t.Id == req.Id {
			todos[i].Status = req.NewStatus
			todos[i].UpdatedAt = time.Now()

			updated = todos[i]

			found = true

			break
		}
	}

	if !found {
		return model.Todo{}, errors.New("no todo with this id")
	}

	err = utils.WriteTodosToFile(todos)

	if err != nil {
		return model.Todo{}, err
	}

	return updated, nil
}

func (t *TodoService) SearchTodo(req dto.SearchTodoRequest) (dto.SearchTodoResponse, error) {
    todos, err := utils.ReadTodosFromFile()
    if err != nil {
        return nil, err
    }

    var results dto.SearchTodoResponse

    keyword := strings.ToLower(strings.TrimSpace(req.Task_Name))

    for _, todo := range todos {
        if strings.Contains(strings.ToLower(todo.Task_name), keyword) {

            results = append(results, dto.SearchTodoItem{
                Id:        todo.Id,
                Task_name: todo.Task_name,
                Priority:  todo.Priority,
                Status:    todo.Status,
            })
        }
    }

    if len(results) == 0 {
        return nil, errors.New("no todo found matching search keyword")
    }

    return results, nil
}

func (t *TodoService) DeleteTodo(id int) error {
	todos, err := utils.ReadTodosFromFile()

	if err != nil {
		return err
	}

	index := -1

	for i, t := range todos {
		if t.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("todo not found")
	}

	todos = append(todos[:index], todos[index+1:]...)

	if err := utils.WriteTodosToFile(todos); err != nil {
		return err
	}

	return nil
} 

func getNextID(todos []model.Todo) int {
    if len(todos) == 0 {
        return 1 // kalo JSON kosong, mulai dari 1
    }

    maxID := 0
    for _, t := range todos {
        if t.Id > maxID {
            maxID = t.Id
        }
    }

    return maxID + 1
}



