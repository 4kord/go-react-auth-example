package users

import (
	"github.com/4kord/go-react-auth/internal/core/domain"
	"database/sql"

	"github.com/4kord/go-react-auth/internal/errs"
)

type Repository interface {
    GetUserById(int) (*domain.User, *errs.Error)
    GetUserByUsername(string) (*domain.User, *errs.Error)
    Authenticate(string, string) (*domain.User, *errs.Error)
    NewUser(string, string) *errs.Error
}

type repository struct {
	DB *sql.DB
}

func New(db *sql.DB) repository {
	return repository{DB: db}
}
