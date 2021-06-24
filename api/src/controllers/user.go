package controllers

import "net/http"

// Create insert an user in the database
func Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
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
