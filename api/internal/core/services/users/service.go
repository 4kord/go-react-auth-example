package users

import (
	"github.com/4kord/go-react-auth/internal/core/repositories/users"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

type Service interface {
	Login(dto.UserRequest) (*dto.UserResponse, *errs.Error)
	Register(dto.UserRequest) *errs.Error
}

type service struct {
	repo users.Repository
}

func New(repo users.Repository) service {
	return service{repo: repo}
}
