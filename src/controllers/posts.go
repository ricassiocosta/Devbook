package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/utils/authentication"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Publish add a new post in the database
func Publish(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(body, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	post.AuthorID = userID

	if err = post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	post.ID, err = repository.Create(post)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, post)
}

// GetPosts return the feed for a given user
func GetPosts(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.GetPosts(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, posts)
}

// GetPost return a single post
func GetPost(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	post, err := repository.GetByID(postID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, post)
}

// UpdatePost updates a post data
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	storedPost, err := repository.GetByID(postID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if storedPost.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errForbidden)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var post models.Post
	if err = json.Unmarshal(body, &post); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = post.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = repository.Update(postID, post); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// DeletePost deletes a post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	userID, err := authentication.GetUserID(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	storedPost, err := repository.GetByID(postID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if storedPost.AuthorID != userID {
		responses.Error(w, http.StatusForbidden, errForbidden)
		return
	}

	if err = repository.Delete(postID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetPostsByUser return all posts from a specific user
func GetPostsByUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	posts, err := repository.GetPostsByUserID(userID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusOK, posts)
}

// Like add one like on post count
func Like(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	if err = repository.Like(postID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// Dislike remove one like on post count
func Dislike(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	postID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewPostsRepository(db)
	if err = repository.Dislike(postID); err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
