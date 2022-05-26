package users

import (
	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
)

func (r repository) AuthenticateX(username, password string) (*domain.User, *errs.Error) {
    user := domain.User{
        Id: 6,
        Username: "test",
        Password: "test",
        Role: "user",
    }
    return &user, nil
}
