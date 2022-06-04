package dto

import (
	"time"

	"github.com/4kord/go-react-auth/internal/errs"
)

type SessionRequest struct {
	RefreshToken string
	Ip           string
}

type SessionResponse struct {
	AccessToken    string    `json:"accessToken"`
	RefreshToken   string    `json:"-"`
	RefreshExpires time.Time `json:"-"`
}

func (r *SessionRequest) Validate() *errs.Error {
	if r.RefreshToken == "" {
		return errs.UnAuthorizedError("No refresh token provided")
	}

	if r.Ip == "" {
		return errs.UnAuthorizedError("No ip prodvided")
	}

	return nil
}
