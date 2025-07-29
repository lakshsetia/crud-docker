package handlers

import (
	"net/http"

	"github.com/lakshsetia/crud-docker/internal/storage"
)

func UserHandler(storage storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
}

func UserByIdHandler(storage storage.Storage) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
	})
}