package sessions

import (
	"database/sql"
	"errors"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
)

func (r repository) NewSessionWithRemoving(session domain.Session) *errs.Error {
	tx, err := r.DB.Begin()
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return errs.ServerError("Error starting db transaction")
	}

	q := "SELECT id FROM sessions WHERE user_id = $1"

	var ids []int
	var id int

	rows, err := tx.Query(q, session.UserId)
	if err != nil {
		rbErr := tx.Rollback()
		if rbErr != nil {
			logger.ErrorLog.Println(err.Error())
			return errs.ServerError("Error rollbacking transaction")
		}
		logger.InfoLog.Println(err.Error())
		return errs.ServerError("Unexpected DB error")
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				break
			} else {
				rbErr := tx.Rollback()
				if rbErr != nil {
					logger.ErrorLog.Println(rbErr.Error())
					return errs.ServerError("Error rollbacking transaction")
				}
				logger.ErrorLog.Println(err.Error())
				return errs.ServerError("Unexpected DB error")
			}
		}

		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		logger.ErrorLog.Println((err.Error()))
		return errs.ServerError("Unexpected DB error")
	}

	if len(ids) >= 5 {
		smallest := ids[0]

		for _, num := range ids[1:] {
			if num < smallest {
				smallest = num
			}
		}

		q := "DELETE FROM sessions WHERE id = $1"

		_, err = tx.Exec(q, smallest)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				logger.ErrorLog.Println(txErr.Error())
				return errs.ServerError("Error rollbacking transcation")
			}
			logger.ErrorLog.Println(err.Error())
			return errs.ServerError("Unexpected DB error")
		}
	}

	q = "INSERT INTO sessions(user_id, refresh_token, ip, expires_at) VALUES($1, $2, $3, $4)"

	_, err = tx.Exec(q, session.UserId, session.RefreshToken, session.Ip, session.ExpiresAt)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			logger.ErrorLog.Println(txErr.Error())
			return errs.ServerError("Error rollbacking transcation")
		}
		logger.ErrorLog.Println(err.Error())
		return errs.ServerError("Unexpected DB error")
	}

	err = tx.Commit()
	if err != nil {
		return errs.ServerError("Error commiting transaction")
	}

	return nil
}
