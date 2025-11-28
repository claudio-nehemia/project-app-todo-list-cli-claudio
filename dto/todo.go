package dto

type CreateTodoRequest struct {
	Task_name	string
	Priority	string
}

type ListTodoResponse struct {
	Id			int
	Task_name 	string
	Priority	string
	Status		string
}

type UpdateTodoRequest struct {
	Id			int
	NewStatus	string
}

type SearchTodoRequest struct {
	Task_Name	string
}

type SearchTodoItem struct {
	Id			int
	Task_name 	string
	Priority	string
	Status		string
}

type SearchTodoResponse []SearchTodoItem