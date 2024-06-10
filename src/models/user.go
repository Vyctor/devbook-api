package models

import (
	"devbook-api/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}
	if err := user.format(stage); err != nil {
		return err
	}
	return nil
}

func (user *User) validate(state string) error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if user.Nick == "" {
		return errors.New("O apelido é obrigatório e não pode estar em branco")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}
	if checkmail.ValidateFormat(user.Email) != nil {
		return errors.New("O email é inválido")
	}
	if user.Password == "" && state == "create" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	if stage == "cadastro" {
		hash, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hash)
	}
	return nil
}
