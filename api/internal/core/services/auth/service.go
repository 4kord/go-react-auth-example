package auth

import (
	"github.com/4kord/go-react-auth/internal/core/repositories/sessions"
	"github.com/4kord/go-react-auth/internal/core/repositories/users"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

type Service interface {
	Login(dto.UserRequest) (*dto.UserResponse, *errs.Error)
	Register(dto.UserRequest) *errs.Error
	Logout(dto.SessionRequest) *errs.Error
	Refresh(dto.SessionRequest) (*dto.UserResponse, *errs.Error)
}

type service struct {
	repo        users.Repository
	sessionRepo sessions.Repository
}

func New(repo users.Repository, sessionRepo sessions.Repository) service {
	return service{
		repo:        repo,
		sessionRepo: sessionRepo,
	}
}
