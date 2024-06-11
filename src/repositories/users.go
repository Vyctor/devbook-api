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

func (repo Users) Follow(userId uint64, followerId uint64) error {
	stmt, err := repo.db.Prepare("INSERT IGNORE INTO seguidores (usuario_id, seguidor_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	if _, err = stmt.Exec(userId, followerId); err != nil {
		return err
	}
	return nil
}

func (repo Users) Unfollow(userId uint64, followerId uint64) error {
	stmt, err := repo.db.Prepare("DELETE FROM seguidores s WHERE s.usuario_id = ? AND s.seguidor_id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {

		}
	}(stmt)
	if _, err = stmt.Exec(userId, followerId); err != nil {
		return err
	}
	return nil
}

func (repo Users) GetFollowers(userId uint64) ([]models.User, error) {
	rows, err := repo.db.Query("SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios u WHERE u.id IN (SELECT s.seguidor_id seguidor FROM seguidores s WHERE s.usuario_id = ?);", userId)
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

func (repo Users) GetFollowing(userId uint64) ([]models.User, error) {
	rows, err := repo.db.Query("SELECT u.id, u.nome, u.nick, u.email, u.criadoEm FROM usuarios u WHERE u.id IN (SELECT s.usuario_id from seguidores s WHERE s.seguidor_id = ?);", userId)
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

func (repo Users) GetPassword(userId uint64) (string, error) {
	row, err := repo.db.Query("SELECT senha FROM usuarios WHERE id = ?", userId)
	if err != nil {
		return "", err
	}
	defer func(row *sql.Rows) {
		_ = row.Close()
	}(row)
	var password string
	if row.Next() {
		if err = row.Scan(&password); err != nil {
			return "", err
		}
	}
	return password, nil
}

func (repo Users) UpdatePassword(userId uint64, password string) error {
	stmt, err := repo.db.Prepare("UPDATE usuarios SET senha = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		_ = stmt.Close()
	}(stmt)
	if _, err = stmt.Exec(password, userId); err != nil {
		return err
	}
	return nil
}
