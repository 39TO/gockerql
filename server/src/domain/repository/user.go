package repository

import "github.com/39TO/gockerql/domain/entity"

type IUserRepository interface {
	RegisterUser(name string, email string, password string) (*entity.User, error)
	FindUserById(id string) (*entity.User, error)
}
