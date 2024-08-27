package repository

import "github.com/39TO/gockerql/domain/entity"

type ITodoRepository interface {
	CreateTodo(title string, user_id string) (*entity.Todo, error)
	DeleteTodo(id string) error
	FindTodoById(id string) (*entity.Todo, error)
	FindTodosByUserId(user_id string) ([]entity.Todo, error)
}
