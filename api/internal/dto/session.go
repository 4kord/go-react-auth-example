package dto

import "github.com/4kord/go-react-auth/internal/errs"

type SessionRequest struct {
    RefreshToken string
    Ip string
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
