package usecase

import (
	"github.com/39TO/gockerql/domain/entity"
	"github.com/39TO/gockerql/domain/repository"
	usecase_error "github.com/39TO/gockerql/error/usecase"
	"golang.org/x/crypto/bcrypt"
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
	if name == "" {
		return nil, usecase_error.NameEmptyError
	}
	if email == "" {
		return nil, usecase_error.EmailEmptyError
	}
	if password == "" {
		return nil, usecase_error.PasswordEmptyError
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user, err := uc.repo.RegisterUser(name, email, string(hashedPassword))
	return user, err
}
