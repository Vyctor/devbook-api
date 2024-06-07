package repositories

import (
	"database/sql"
	"devbook-api/src/models"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (uint64, error) {
	stmt, err := repo.db.Prepare("INSERT INTO devbook.usuarios (nome , nick, email, senha) VALUES (?, ?, ?, ?)")

	if err != nil {
		return 0, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err)
		}
	}(stmt)

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)

	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return uint64(lastId), nil
}
