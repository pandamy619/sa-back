package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	name, err := ioutil.TempDir("", "tmp")
	if err != nil {
		fmt.Println(err)
	}
	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile(name, "upload-*.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
}
