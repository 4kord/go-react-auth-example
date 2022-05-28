package auth

import (
	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

func (s service) Register(r dto.UserRequest) *errs.Error {
	err := r.Validate()
	if err != nil {
		return err
	}

	err = s.repo.NewUser(domain.User{
        Username: r.Username,
        Password: r.Password,
    })
	if err != nil {
		return err
	}

	return nil
}
