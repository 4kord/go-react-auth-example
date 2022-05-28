package utils

import (
	"database/sql"

	"github.com/4kord/go-react-auth/internal/logger"
)

func ExecTx(db *sql.DB, fn func(*sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return err
	}

	err = fn(tx)
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		logger.ErrorLog.Println(err.Error())
		return err
	}

	return nil
}
