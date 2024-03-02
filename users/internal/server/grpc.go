package server

import (
	"github.com/Prokopevs/kanban/schema"
	"github.com/Prokopevs/kanban/users/internal/core"
)

type GRPC struct {
	schema.UnimplementedUsersServer

	service core.Service
}

func NewGRPC(usersService core.Service) *GRPC {
	return &GRPC{
		service: usersService,
	}
}
