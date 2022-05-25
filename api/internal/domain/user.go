package domain

import "github.com/4kord/go-react-auth/internal/errs"

type User struct {
    Id int
    Username string
    Password string
    Role string
}

type UserRepository interface {
    GetUserById(int) (*User, *errs.Error)
    GetUserByUsername(string) (*User, *errs.Error)
    NewUser(string, string) *errs.Error
}
