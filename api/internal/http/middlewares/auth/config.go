package auth

import "github.com/4kord/go-react-auth/internal/core/repositories/users"

const (
    AdminRole = "admin"
    UserRole = "user"
)

type Config struct {
    Repo users.Repository
    Role string
}

