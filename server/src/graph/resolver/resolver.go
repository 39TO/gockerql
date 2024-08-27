package resolver

import "github.com/39TO/gockerql/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ucTodo usecase.ITodoUsecase
	ucUser usecase.IUserUsecase
}

func NewResolver(ucTodo usecase.ITodoUsecase, ucUser usecase.IUserUsecase) Resolver {
	return Resolver{
		ucTodo: ucTodo,
		ucUser: ucUser,
	}
}
