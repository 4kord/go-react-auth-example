package sessions

import (
	"database/sql"
	"errors"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/lib/pq"
)

func (r repository) GetSession(refreshToken string) (*domain.Session, *errs.Error) {
    var session domain.Session

    q := "SELECT * FROM sessions WHERE refresh_token = $1"

    row := r.DB.QueryRow(q, refreshToken)

    err := row.Scan(&session.Id, &session.UserId, &session.RefreshToken, &session.Ip, &session.ExpiresAt, &session.CreatedAt)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        if errors.Is(err, sql.ErrNoRows) {
            return nil, errs.UnAuthorizedError("Session no longer exists")
        } else if err, ok := err.(*pq.Error); ok && err.Code == "22P02" {
            return nil, errs.UnAuthorizedError("Session no longer exists")
        }
        return nil, errs.ServerError("Unexpected DB error")
    }

    return &session, nil
}
