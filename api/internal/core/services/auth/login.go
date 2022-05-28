package auth

import (
	"os"
	"strconv"
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

    rtExp, e := strconv.Atoi(os.Getenv("RT_EXP"))
    if e != nil {
        return nil, errs.ServerError("Error convertint RT_EXP")
    }
    
    rtExpAt := time.Now().UTC().Add(time.Duration(rtExp) * time.Minute)

	err = s.sessionRepo.NewSessionWithRemoving(domain.Session{
		UserId:       user.Id,
		RefreshToken: rT,
		Ip:           r.Ip,
		ExpiresAt:    rtExpAt,
	})
	if err != nil {
		return nil, err
	}

	userResponse := &dto.UserResponse{
		Id:           user.Id,
		Username:     user.Username,
		Role:         user.Role,
		AccessToken:  aT,
		RefreshToken: rT,
        RefreshExpires: rtExpAt,
	}

	return userResponse, nil
}
