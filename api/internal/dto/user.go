package dto

import (
	"net/http"

	"github.com/4kord/go-react-auth/internal/errs"
)

type UserRequest struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserResponse struct {
    Id int `json:"id"`
    Username string `json:"username"`
    Password string `json:"password"`
    Role string `json:"role"`
    AccessToken string `json:"accessToken"`
}

func (r UserRequest) Validate() *errs.Error {
    if r.Username == "" || r.Password == "" {
        return &errs.Error{
            Code:    http.StatusUnprocessableEntity,
            Message: "Make sure fields are not empty",
        }
    }
    return nil
}
