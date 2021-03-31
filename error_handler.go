package main

import (
	"encoding/json"
	"net/http"
)

// Error entity
type Error struct {
	Code      string `json:"code"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}

// NewError generates an error in json format.
func NewError(w http.ResponseWriter, error string, code int) {
	err := Error{Code: error, Message: error, Timestamp: getTimestamp()}
	errorJson(w, err, code)
}

// ErrorJson replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be json.
func errorJson(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(err)
}
