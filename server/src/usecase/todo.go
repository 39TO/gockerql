package usecase

import (
	"github.com/39TO/gockerql/domain/entity"
	"github.com/39TO/gockerql/domain/repository"
	usecase_error "github.com/39TO/gockerql/error/usecase"
)

var _ ITodoUsecase = &TodoUsecase{}

type TodoUsecase struct {
	repo repository.ITodoRepository
}

type ITodoUsecase interface {
	CreateTodo(title string, user_id string) (*entity.Todo, error)
	DeleteTodo(id string) error
}

func NewTodoUsecase(repo repository.ITodoRepository) ITodoUsecase {
	return &TodoUsecase{
		repo: repo,
	}
}

func (uc *TodoUsecase) CreateTodo(title string, user_id string) (*entity.Todo, error) {
	if title == "" {
		return nil, usecase_error.TitleEmptyError
	}
	if user_id == "" {
		return nil, usecase_error.IdEmptyError
	}

	todo, err := uc.repo.CreateTodo(title, user_id)
	return todo, err
}

func (usecase *TodoUsecase) DeleteTodo(id string) error {
	if id == "" {
		return usecase_error.IdEmptyError
	}

	err := usecase.repo.DeleteTodo(id)
	return err
}
