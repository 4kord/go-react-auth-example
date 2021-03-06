package sessions

import "github.com/4kord/go-react-auth/internal/errs"

func (r repository) DeleteSession(rT string) *errs.Error {
	q := "DELETE FROM sessions WHERE refresh_token = $1"

	_, err := r.DB.Exec(q, rT)
	if err != nil {
		return errs.ServerError("Unexpected DB error")
	}

	return nil
}
