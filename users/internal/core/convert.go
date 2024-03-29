package core

import "github.com/Prokopevs/kanban/users/internal/pg"

func (n *User) toDB() *pg.User {
	return &pg.User{
		ID:             n.ID,
		Name:           n.Name,
		Email:          n.Email,
		EmailConfirmed: n.EmailConfirmed,
		PasswordHash:   n.PasswordHash,
	}
}

func convertDBUserToService(user *pg.User) *User {
	return &User{
		ID:             user.ID,
		Name:           user.Name,
		Email:          user.Email,
		EmailConfirmed: user.EmailConfirmed,
		PasswordHash:   user.PasswordHash,
	}
}
