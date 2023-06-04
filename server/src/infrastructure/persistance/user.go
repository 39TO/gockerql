package persistance

import (
	"database/sql"

	"github.com/39TO/gockerql/domain/entity"
	"github.com/39TO/gockerql/domain/repository"
)

var _ repository.IUserRepository = &userRepository{}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (repo *userRepository) RegisterUser(name string, email string, password string) (*entity.User, error) {
	statement := "INSERT INTO users (name, email, password) VALUES (?, ?, ?)"

	stmt, err := repo.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user := &entity.User{}
	res, err := stmt.Exec(name, email, password)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}

	user.Id = string(id)
	user.Name = name
	user.Email = email
	user.Password = password

	return user, nil
}
