package auth

import (
	"github.com/4kord/go-react-auth/internal/dto"
	"github.com/4kord/go-react-auth/internal/errs"
)

func (s service) Logout(r dto.SessionRequest) *errs.Error {
	return s.sessionRepo.DeleteSession(r.RefreshToken)
}
