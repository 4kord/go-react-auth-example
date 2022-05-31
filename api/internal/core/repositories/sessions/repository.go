package sessions

import (
	"database/sql"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
)

type Repository interface {
	GetSession(string) (*domain.Session, *errs.Error)
	NewSessionWithRemoving(domain.Session) *errs.Error
	DeleteSession(string) *errs.Error
	DeleteSessionById(int) *errs.Error
}

type repository struct {
	DB *sql.DB
}

func New(db *sql.DB) repository {
	return repository{DB: db}
}
