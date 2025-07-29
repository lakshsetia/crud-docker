package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/lakshsetia/crud-docker/internal/storage"
	"github.com/lakshsetia/crud-docker/internal/types"
	"github.com/lakshsetia/crud-docker/internal/utils/error"
	"github.com/lakshsetia/crud-docker/internal/utils/json"
)

func UserHandler(storage storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.Method {
		case http.MethodGet:
			users, err := storage.GetUsers()
			if err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelDatabase, err.Error()), http.StatusInternalServerError)
				return
			}
			json.WriteJSON(w, users, http.StatusOK)
		case http.MethodPost:
			var user types.User
			json.ReadJSON(w, r, &user)
			if err := user.Validate(); err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageInvalidJSONRequest), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
			if err := storage.CreateUser(user); err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelDatabase, err.Error()), http.StatusInternalServerError)
				return
			}
		default:
			json.WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageInvalidMethod), http.StatusMethodNotAllowed)
			return
		}
	})
}

func UserByIdHandler(storage storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		pathSegments := strings.Split(r.URL.Path, "/")
		if len(pathSegments) != 3 {
			json.WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageBadURL), http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			json.WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageBadURL), http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			user, err := storage.GetUserById(id)
			if err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelDatabase, err.Error()), http.StatusInternalServerError)
				return
			}
			json.WriteJSON(w, user, http.StatusOK)
		case http.MethodPut:
			var newUser types.User
			json.ReadJSON(w, r, &newUser)
			if err := newUser.Validate(); err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageInvalidJSONRequest), http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			if err := storage.UpdateUserById(id, newUser); err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelDatabase, err.Error()), http.StatusInternalServerError)
				return
			}
		case http.MethodDelete:
			if err := storage.DeleteUserById(id); err != nil {
				json.WriteJSON(w, error.NewErrorResponse(error.LevelDatabase, err.Error()), http.StatusInternalServerError)
			}
		default:
			json.WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageInvalidMethod), http.StatusMethodNotAllowed)
			return
		}
	})
}