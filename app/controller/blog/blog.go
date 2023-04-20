package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/snretuerma/gorest/model"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts, err := model.GetAllPosts()

	if err != nil {
		writeError(w, err)
		return
	} else {
		json.NewEncoder(w).Encode(posts)
	}
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam := chi.URLParam(r, "id")

	post, err := model.GetPost(idParam)

	if err != nil {
		writeError(w, err)
		return
	} else {
		json.NewEncoder(w).Encode(post)
	}
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var post model.Post

	err := decoder.Decode(&post)

	if err != nil {
		writeError(w, err)
		return
	}

	err = model.CreatePost(post)

	if err != nil {
		writeError(w, err)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam, err := strconv.ParseUint(string(chi.URLParam(r, "id")), 10, 64)

	if err != nil {
		writeError(w, err)
		return
	}

	decoder := json.NewDecoder(r.Body)

	var post model.Post

	err = decoder.Decode(&post)

	if err != nil {
		writeError(w, err)
		return
	}

	err = model.UpdatePost(idParam, post)

	if err != nil {
		writeError(w, err)
		return
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idParam, err := strconv.ParseUint(string(chi.URLParam(r, "id")), 10, 64)

	if err != nil {
		writeError(w, err)
		return
	}

	err = model.DeletePost(idParam)

	if err != nil {
		writeError(w, err)
		return
	} else {
		json.NewEncoder(w).Encode(http.StatusOK)
	}

}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}
