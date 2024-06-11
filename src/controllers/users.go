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
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare("cadastro"); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	repo := repositories.NewUserRepository(db)
	userId, err := repo.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	user.ID = userId
	responses.JSON(w, http.StatusCreated, user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, err)
		}
	}(db)
	repo := repositories.NewUserRepository(db)
	users, err := repo.Get(nameOrNick)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/users/"), 10, 64)
	if err != nil {
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
	user, err := repo.GetById(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/users/"), 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	userIdToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	if userId != userIdToken {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	tokenUser, err := authentication.ExtractUserId(r)
	if userId != tokenUser {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var user = models.User{}
	user.ID = userId
	if err = json.Unmarshal(reqBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if err = user.Prepare("edicao"); err != nil {
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
	if err = repo.Update(user); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.ParseUint(strings.TrimPrefix(r.URL.Path, "/users/"), 10, 64)
	if err != nil {
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
	if err = repo.Delete(userId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
}

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followerId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if followerId == userId {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepository(db)
	if err = repo.Follow(userId, followerId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	unfollowerId, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if unfollowerId == userId {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepository(db)
	if err = repo.Unfollow(userId, unfollowerId); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepository(db)
	followers, err := repo.GetFollowers(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepository(db)
	followers, err := repo.GetFollowing(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, followers)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userToken, err := authentication.ExtractUserId(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	userId, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	if userToken != userId {
		responses.Error(w, http.StatusForbidden, err)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}
	var password models.ChangePassword
	if err = json.Unmarshal(body, &password); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}
	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	repo := repositories.NewUserRepository(db)
	passwordOnDb, err := repo.GetPassword(userId)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = security.VerifyPassword(passwordOnDb, password.OldPassword); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}
	hashedPassword, err := security.Hash(password.NewPassword)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	if err = repo.UpdatePassword(userId, string(hashedPassword)); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusNoContent, nil)
}
