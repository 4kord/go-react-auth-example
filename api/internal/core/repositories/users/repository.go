package users

import (
	"database/sql"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
)

type Repository interface {
    GetUser(int) (*domain.User, *errs.Error)
    Authenticate(string, string) (*domain.User, *errs.Error)
    NewUser(domain.User) *errs.Error
}

type repository struct {
	DB *sql.DB
}

func New(db *sql.DB) repository {
	return repository{DB: db}
}
