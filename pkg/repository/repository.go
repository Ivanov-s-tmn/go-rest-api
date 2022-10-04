package repository

import (
	"github.com/ivanov-s-tmn/go-rest-api"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GetUser(username, password string) (rest.User, error)
}

type TodoList interface {
	Create(userId int, list rest.TodoList) (int, error)
	GetAll(userId int) ([]rest.TodoList, error)
	GetById(userId, listId int) (rest.TodoList, error)
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
	}
}
