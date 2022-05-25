package users

import (
	"github.com/4kord/go-react-auth/internal/core/repositories/usersrepo"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

type Service interface {
	Login(dto.UserRequest) (*dto.UserResponse, *errs.Error)
	Register(dto.UserRequest) *errs.Error
}

type service struct {
	Repo usersrepo.Repository
}

func New(repo usersrepo.Repository) service {
	return service{Repo: repo}
}
