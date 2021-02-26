package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	confFile     = "config"
	pathConfFile = "."
	typeConfFile = "yml"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Parse our multipart form, 10 << 20 specifies a maximum
	// upload of 10 MB files.
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create a temporary file within our temp-images directory that follows
	// a particular naming pattern
	tempFile, err := ioutil.TempFile("temp-report", "upload-*.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
}

func main() {
	config, err := loadEnv(confFile, pathConfFile, typeConfFile)
	if err != nil {
		fmt.Println(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", uploadHandler)

	if err := http.ListenAndServe(":"+config.Port, mux); err != nil {
		log.Fatal(err)
	}

}
