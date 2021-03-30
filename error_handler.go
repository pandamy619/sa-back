package main

import (
	"encoding/json"
	"net/http"
)

// Error entity
type Error struct {
	Err         error  `json:"error"`
	Description string `json:"description"`
	Code        int    `json:"code"`
}

// ErrorJson replies to the request with the specified error message and HTTP code.
// It does not otherwise end the request; the caller should ensure no further
// writes are done to w.
// The error message should be json.
func ErrorJson(w http.ResponseWriter, err interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(err)
}
