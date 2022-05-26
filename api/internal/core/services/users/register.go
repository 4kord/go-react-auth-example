package users

import (
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

func (s service) Register(u dto.UserRequest) *errs.Error {
	err := u.Validate()
	if err != nil {
		return err
	}

	err = s.repo.NewUser(u.Username, u.Password)
	if err != nil {
		return err
	}

	return nil
}
