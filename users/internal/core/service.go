package core

import (
	"context"

	"github.com/Prokopevs/kanban/users/internal/pg"
)

var _ Service = (*ServiceImpl)(nil)

type Service interface {
	AddUser(context.Context, *User) error
	GetUser(context.Context, string) (*User, bool, error)
	GetUserByEmail(context.Context, string) (*User, bool, error)
	IsUserWithEmailExists(context.Context, string) (bool, error)
	IsValidUserCredentials(context.Context, string, string) (bool, error)
	UpdateUser(context.Context, *User) error
}

type ServiceImpl struct {
	db pg.DB
}

func NewServiceImpl(db pg.DB) *ServiceImpl {
	return &ServiceImpl{
		db: db,
	}
}
