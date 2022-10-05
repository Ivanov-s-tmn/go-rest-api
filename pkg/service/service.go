package service

import (
	"github.com/ivanov-s-tmn/go-rest-api"
	"github.com/ivanov-s-tmn/go-rest-api/pkg/repository"
)

type Authorization interface {
	CreateUser(user rest.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list rest.TodoList) (int, error)
	GetAll(userId int) ([]rest.TodoList, error)
	GetById(userId, listId int) (rest.TodoList, error)
	Update(userId, listId int, input rest.UpdateListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId, listId int, item rest.TodoItem) (int, error)
	GetAll(userId, listId int) ([]rest.TodoItem, error)
	GetById(userId, itemId int) (rest.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input rest.UpdateItemInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		TodoList:      NewTodoListService(repo.TodoList),
		TodoItem:      NewTodoItemService(repo.TodoItem, repo.TodoList),
	}
}
