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
	Update(userId, listId int, input rest.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(listId int, item rest.TodoItem) (int, error)
	GetAll(userId, listId int) ([]rest.TodoItem, error)
	GetById(userId, itemId int) (rest.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input rest.UpdateItemInput) error
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
		TodoItem:      NewTodoItemPostgres(db),
	}
}
