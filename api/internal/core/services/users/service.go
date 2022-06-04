package users

import (
	"github.com/4kord/go-react-auth/internal/core/repositories/users"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

type Service interface {
	GetUser(id int) (*dto.UserResponse, *errs.Error)
}

type service struct {
	repo users.Repository
}

func New(repo users.Repository) service {
	return service{
		repo: repo,
	}
}
