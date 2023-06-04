package usecase

import (
	"github.com/39TO/gockerql/domain/entity"
	"github.com/39TO/gockerql/domain/repository"
)

var _ IUserUsecase = &UserUsecase{}

type UserUsecase struct {
	repo repository.IUserRepository
}

type IUserUsecase interface {
	RegisterUser(name string, email string, password string) (*entity.User, error)
}

func NewUserUsecase(repo repository.IUserRepository) IUserUsecase {
	return &UserUsecase{
		repo: repo,
	}
}

func (uc *UserUsecase) RegisterUser(name string, email string, password string) (*entity.User, error) {
	user, err := uc.repo.RegisterUser(name, email, password)
	return user, err
}
