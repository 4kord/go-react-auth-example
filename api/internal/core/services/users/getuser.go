package users

import (
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

func (s service) GetUser(id int) (*dto.UserResponse, *errs.Error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		Id:       user.Id,
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
