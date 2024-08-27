package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"fmt"

	"github.com/39TO/gockerql/graph"
	"github.com/39TO/gockerql/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodoInput) (*model.Todo, error) {
	todo, err := r.ucTodo.CreateTodo(input.Title, input.UserID)

	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:    todo.Id,
		Title: todo.Title,
		Done:  todo.Done,
	}, nil
}

// DeleteTodo is the resolver for the deleteTodo field.
func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.DeleteTodoInput) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteTodo - deleteTodo"))
}

// RegisterUser is the resolver for the registerUser field.
func (r *mutationResolver) RegisterUser(ctx context.Context, input model.NewUserInput) (*model.User, error) {
	user, err := r.ucUser.RegisterUser(input.Name, input.Email, input.Password)

	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
