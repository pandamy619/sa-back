package main

import "net/http"

func handlers(mux *http.ServeMux) {
	mux.HandleFunc("/upload", uploadHandler)
}
