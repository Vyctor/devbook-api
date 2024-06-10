package repositories

import (
	"database/sql"
	"devbook-api/src/models"
	"fmt"
)

type Users struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Users {
	return &Users{db}
}

func (repo Users) Create(user models.User) (uint64, error) {
	stmt, err := repo.db.Prepare("INSERT INTO usuarios (nome , nick, email, senha) VALUES (?, ?, ?, ?)")
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

func (repo Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)
	rows, err := repo.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios  WHERE nick LIKE ? OR nome LIKE ?", nameOrNick, nameOrNick)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (repo Users) GetById(userId uint64) (models.User, error) {
	row, err := repo.db.Query("SELECT id, nome, nick, email, criadoEm FROM usuarios WHERE id = ?", userId)
	if err != nil {
		return models.User{}, err
	}
	defer func(row *sql.Rows) {
		_ = row.Close()
	}(row)
	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (repo Users) Update(user models.User) error {
	stmt, err := repo.db.Prepare("UPDATE usuarios SET nome = ?, nick = ?, email = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	if _, err = stmt.Exec(user.Name, user.Nick, user.Email, user.ID); err != nil {
		return err
	}
	return nil
}

func (repo Users) Delete(userId uint64) error {
	stmt, err := repo.db.Prepare("DELETE FROM usuarios WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	if _, err = stmt.Exec(userId); err != nil {
		return err
	}
	return nil
}

func (repo Users) FindByEmail(email string) (models.User, error) {
	row, err := repo.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)

	if err != nil {
		return models.User{}, err
	}

	defer func(row *sql.Rows) {
		_ = row.Close()
	}(row)
	var user models.User
	if row.Next() {
		if err = row.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}
