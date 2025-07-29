package json

import (
	"encoding/json"
	"net/http"

	"github.com/lakshsetia/crud-docker/internal/utils/error"
)

func WriteJSON(w http.ResponseWriter, response any, statusCode int) {
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func ReadJSON(w http.ResponseWriter, r *http.Request, request any) {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		WriteJSON(w, error.NewErrorResponse(error.LevelBackend, error.MessageBadRequest), http.StatusBadRequest)
	}
}