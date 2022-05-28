package sessions

import (
	"database/sql"
	"errors"

	"github.com/4kord/go-react-auth/internal/core/domain"
	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/utils"
)

func (r repository) NewSessionWithRemoving(session domain.Session) *errs.Error {
	err := utils.ExecTx(r.DB, func(tx *sql.Tx) error {
		q := "SELECT id FROM sessions WHERE user_id = $1"

		var ids []int
		var id int

		rows, err := tx.Query(q, session.UserId)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			err := rows.Scan(&id)
			if err != nil {
				if errors.Is(err, sql.ErrNoRows) {
					break
				}
				return err
			}

			ids = append(ids, id)
		}

		if err = rows.Err(); err != nil {
			return err
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
				return err
			}
		}

		q = "INSERT INTO sessions(user_id, refresh_token, ip, expires_at) VALUES($1, $2, $3, $4)"

		_, err = tx.Exec(q, session.UserId, session.RefreshToken, session.Ip, session.ExpiresAt)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return errs.ServerError("Unexpected db error")
	}

	return nil
}
