package users

import (
	"github.com/4kord/go-react-auth/internal/errs"
	"github.com/4kord/go-react-auth/internal/logger"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func (r repository) NewUser(username, password string) *errs.Error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        logger.ErrorLog.Println(err.Error())
        return errs.ServerError("Unexpected bcrypt error")
    }

    q := "INSERT INTO users(username, password) VALUES($1, $2)"
    
    _, err = r.DB.Exec(q, username, hashedPassword)
    if err, ok := err.(*pq.Error); ok {
        logger.ErrorLog.Println(err.Error())
        if err.Constraint == "uc_users_username" {
            return errs.ConflictError("Username is already taken")
        } else {
            return errs.ServerError("Unexpected DB error")
        }
    }

    return nil
}
