package persistance

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/39TO/gockerql/domain/entity"
	"github.com/39TO/gockerql/domain/repository"
	db_error "github.com/39TO/gockerql/error/infrastructure"
)

var _ repository.ITodoRepository = &TodoRepository{}

type TodoRepository struct {
	db *sql.DB
}

func NewTodoRepository(db *sql.DB) *TodoRepository {
	return &TodoRepository{
		db: db,
	}
}

func (repo *TodoRepository) CreateTodo(title string, user_id string) (*entity.Todo, error) {
	statement := "INSERT INTO todos (title, user_id) VALUES (?, ?)"

	stmt, err := repo.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v : %v", db_error.StatementError, err)
	}
	defer stmt.Close()

	todo := &entity.Todo{}
	res, err := stmt.Exec(title, user_id)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v : %v", db_error.ExecError, err)
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, fmt.Errorf("%v : %v", db_error.LastInsertError, err)
	}

	todo.Id = string(id)
	todo.Title = title
	todo.Done = false

	return todo, nil
}

func (repo *TodoRepository) DeleteTodo(id string) error {
	statement := "DELETE FROM todos WHERE id = ?"

	stmt, err := repo.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("%v : %v", db_error.StatementError, err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	if err != nil {
		log.Println(err)
		return fmt.Errorf("%v : %v", db_error.ExecError, err)
	}

	return nil
}

func (repo *TodoRepository) FindTodoById(id string) (*entity.Todo, error) {
	statement := "SELECT id, title, done FROM todos WHERE id = ?"

	stmt, err := repo.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v : %v", db_error.StatementError, err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	todo := &entity.Todo{}
	err = row.Scan(&todo.Id, &todo.Title, &todo.Done)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v : %v", db_error.ExecError, err)
	}

	return todo, nil
}

func (repo *TodoRepository) FindTodosByUserId(user_id string) ([]entity.Todo, error) {
	statement := "SELECT id, title, done FROM todos WHERE user_id = ?"

	stmt, err := repo.db.Prepare(statement)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v : %v", db_error.StatementError, err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(user_id)

	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("%v : %v", db_error.ExecError, err)
	}
	defer rows.Close()

	todos := []entity.Todo{}
	for rows.Next() {
		todo := entity.Todo{}
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Done)
		if err != nil {
			log.Println(err)
			return nil, fmt.Errorf("%v : %v", db_error.ExecError, err)
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
