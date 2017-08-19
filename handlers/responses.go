package handlers

import (
	"encoding/json"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, body interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if body != nil {
		_ = json.NewEncoder(w).Encode(body)
	}
}

func responseOK(w http.ResponseWriter, body interface{}) {
	writeJSON(w, http.StatusOK, body)
}

func responseCreated(w http.ResponseWriter, body interface{}, location string) {
	w.Header().Set("Location", location)
	writeJSON(w, http.StatusCreated, body)
}

func responseNoContent(w http.ResponseWriter) {
	writeJSON(w, http.StatusNoContent, &errorMessage{http.StatusText(http.StatusNoContent)})
}

func responseNotFound(w http.ResponseWriter) {
	writeJSON(w, http.StatusNotFound, &errorMessage{http.StatusText(http.StatusNotFound)})
}

func responseInternalServerError(w http.ResponseWriter, err error) {
	writeJSON(w, http.StatusInternalServerError, &errorMessage{err.Error()})
}
