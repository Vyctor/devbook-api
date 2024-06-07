package database

import (
	"database/sql"
	"devbook-api/src/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.DbSringConnection)

	fmt.Println(config.DbSringConnection)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
