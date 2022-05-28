package auth

import (
	"time"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/google/uuid"
)

func (s service) Login(r dto.UserRequest) (*dto.UserResponse, *errs.Error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}

	user, err := s.repo.Authenticate(r.Username, r.Password)
	if err != nil {
		return nil, err
	}

	aT, err := user.GenerateToken()
	if err != nil {
		return nil, err
	}

	rT := uuid.New().String()

	err = s.sessionRepo.NewSessionWithRemoving(domain.Session{
		UserId:       user.Id,
		RefreshToken: rT,
		Ip:           r.Ip,
		ExpiresAt:    time.Now().UTC().Add(24 * time.Hour),
	})
	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Password:     user.Password,
		Role:         user.Role,
		AccessToken:  aT,
		RefreshToken: rT,
	}

	return userResponse, nil
}
