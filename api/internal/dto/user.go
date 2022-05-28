package dto

import (
	"net/http"
	"time"

	"github.com/4kord/go-react-auth/internal/errs"
)

type UserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Ip string `json:"-"`
}

type UserResponse struct {
    Id int `json:"id"`
    Username string `json:"username"`
    Role string `json:"role"`
    AccessToken string `json:"accessToken"`
    RefreshToken string `json:"-"`
    RefreshExpires time.Time `json:"-"`
}

func (r *UserRequest) Validate() *errs.Error {
    if r.Username == "" || r.Password == "" {
        return &errs.Error{
            Code:    http.StatusUnprocessableEntity,
            Message: "Make sure fields are not empty",
        }
    }
    return nil
}
