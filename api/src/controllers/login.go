package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils/security"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

var (
	errWrongEmailOrPassword = errors.New("email or password is wrong")
)

// Login is responsible for authenticate an user
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Error(w, http.StatusBadGateway, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepositories(db)
	storedUser, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.Compare(storedUser.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, errWrongEmailOrPassword)
		return
	}

	w.Write([]byte("logged in!"))
}
