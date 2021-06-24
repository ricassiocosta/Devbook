package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Create an user in the app
func Create(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewUsersRepositories(db)
	repository.Create(user)
}

// Index return all registered users in the database
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

// Show return an specific user
func Show(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting an user"))
}

// Update an specific user
func Update(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user"))
}

// Delete an specific user
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting user"))
}
