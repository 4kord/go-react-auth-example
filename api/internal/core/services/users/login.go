package users

import (
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

func (s service) Login(u dto.UserRequest) (*dto.UserResponse, *errs.Error) {
    err := u.Validate()
    if err != nil {
        return nil, err
    }

    user, err := s.repo.Authenticate(u.Username, u.Password)
    if err != nil {
        return nil, err
    }

    aT, err := user.GenerateToken()
    if err != nil {
        return nil, err
    }

    userResponse := &dto.UserResponse{
        Id: user.Id,
        Username: user.Username,
        Password: user.Password,
        Role: user.Role,
        AccessToken: aT,
    }

    return userResponse, nil
}
