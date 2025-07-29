package storage

import "github.com/lakshsetia/crud-docker/internal/types"

type Storage interface {
	GetUsers() ([]types.User, error)
	CreateUser(user types.User) error
	GetUserById(id int) (types.User, error)
	UpdateUserById(id int, newUser types.User) error
	DeleteUserById(id int) error
}