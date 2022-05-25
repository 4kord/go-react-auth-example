package domain

import (
	"database/sql"
	"errors"

	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"golang.org/x/crypto/bcrypt"
)

type UserRepositoryAdapter struct {
	DB *sql.DB
}

func NewUserRepositoryAdapter(db *sql.DB) UserRepositoryAdapter {
	return UserRepositoryAdapter{DB: db}
}

func (a UserRepositoryAdapter) GetUserById(id int) (*User, *errs.Error) {
    var user User

	q := "SELECT * FROM users WHERE id = $1"

    row := a.DB.QueryRow(q, id)

    err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errs.UnAuthorizedError("Error no rows")
        } else {
            return nil, errs.ServerError("Unexpected DB error")
        }
    }

    return &user, nil
}

func (a UserRepositoryAdapter) GetUserByUsername(username string) (*User, *errs.Error) {
    var user User

	q := "SELECT * FROM users WHERE username = $1"

    row := a.DB.QueryRow(q, username)

    err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errs.UnAuthorizedError("Error no rows")
        } else {
            return nil, errs.ServerError("Unexpected DB error")
        }
    }

    return &user, nil
}

func (a UserRepositoryAdapter) NewUser(username, password string) *errs.Error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        return errs.ServerError("Unexpected bcrypt error")
    }

    q := "INSERT INTO users(username, password) VALUES($1, $2)"
    
    _, err = a.DB.Exec(q, username, hashedPassword)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        return errs.ServerError("Unexpected DB error")
    }

    return nil
}
