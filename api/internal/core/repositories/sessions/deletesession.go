package sessions

import "github.com/4kord/go-react-auth/internal/errs"

func (r repository) DeleteSession(id int) *errs.Error {
    q := "DELETE FROM sessions WHERE id = $1"

    _, err := r.DB.Exec(q, id)
    if err != nil {
        return errs.ServerError("Unexpected DB error")
    }

    return nil
}
