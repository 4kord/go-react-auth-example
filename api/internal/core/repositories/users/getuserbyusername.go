package users

import (
	"database/sql"
	"errors"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
)

func (r repository) GetUserByUsername(username string) (*domain.User, *errs.Error) {
    var user domain.User

	q := "SELECT * FROM users WHERE username = $1"

    row := r.DB.QueryRow(q, username)

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
