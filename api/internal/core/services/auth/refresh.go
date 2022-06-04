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

func (s service) Refresh(r dto.SessionRequest) (*dto.SessionResponse, *errs.Error) {
	err := r.Validate()
	if err != nil {
		return nil, err
	}

	session, err := s.sessionRepo.GetSession(r.RefreshToken)
	if err != nil {
		return nil, err
	}

	if valide := session.ValidateExpiry(); !valide {
		err = s.sessionRepo.DeleteSessionById(session.Id)
		if err != nil {
			return nil, err
		}

		return nil, errs.UnAuthorizedError("Session has expired")
	}

	if valid := session.ValidateIp(r.Ip); !valid {
		err = s.sessionRepo.DeleteSessionById(session.Id)
		if err != nil {
			return nil, err
		}

		return nil, errs.UnAuthorizedError("Ip conflict")
	}

	aT, err := session.GenerateToken()
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
		UserId:       session.UserId,
		RefreshToken: rT,
		Ip:           r.Ip,
		ExpiresAt:    rtExpAt,
	})

	if err != nil {
		return nil, err
	}

	sessionResponse := &dto.SessionResponse{
		AccessToken:    aT,
		RefreshToken:   rT,
		RefreshExpires: rtExpAt,
	}

	return sessionResponse, nil
}
