package service

import (
	"github.com/4kord/go-react-auth/internal/domain"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

type UserService interface {
	Login(dto.UserRequest) (*dto.UserResponse, *errs.Error)
	Register(dto.UserRequest) *errs.Error
}

type DefaultUserService struct {
	Repo domain.UserRepository
}

func NewDefaultUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{Repo: repo}
}

func (s DefaultUserService) Register(u dto.UserRequest) *errs.Error {
	err := u.Validate()
	if err != nil {
		return err
	}

	err = s.Repo.NewUser(u.Username, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (s DefaultUserService) Login(u dto.UserRequest) (*dto.UserResponse, *errs.Error) {
	return nil, nil
}
