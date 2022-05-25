package db

import (
	"database/sql"
)

func Setup(driver, dsn string) (*sql.DB, error) {
	dbConn, err := sql.Open(driver, dsn)
	if err != nil {
        return nil, err
	}

	err = dbConn.Ping()
	if err != nil {
        return nil, err
	}

    return dbConn, nil
}
