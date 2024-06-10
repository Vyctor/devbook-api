package controllers

import (
	"database/sql"
	"devbook-api/authentication"
	"devbook-api/src/database"
	"devbook-api/src/models"
	"devbook-api/src/repositories"
	"devbook-api/src/responses"
	"devbook-api/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user = models.User{}
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	repo := repositories.NewUserRepository(db)
	dbUserData, err := repo.FindByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	err = security.VerifyPassword(dbUserData.Password, string(user.Password))
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	token, _ := authentication.CreateToken(dbUserData.ID)

	responses.JSON(w, http.StatusOK, token)
}
