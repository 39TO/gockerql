package resolver

import "github.com/39TO/gockerql/usecase"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Todo *TodoResolver
	User *UserResolver
}

type TodoResolver struct {
	uc usecase.ITodoUsecase
}

type UserResolver struct {
	uc usecase.IUserUsecase
}

func NewTodoResolver(uc usecase.ITodoUsecase) *TodoResolver {
	return &TodoResolver{
		uc: uc,
	}
}

func NewUserResolver(uc usecase.IUserUsecase) *UserResolver {
	return &UserResolver{
		uc: uc,
	}
}
