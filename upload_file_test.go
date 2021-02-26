package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestUpload(t *testing.T) {
	path := "./config.yaml"
	file, err := os.Open(path)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(file)

	defer file.Close()
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		writer.Close()
		t.Error(err)
	}
	io.Copy(part, file)
	writer.Close()

	req := httptest.NewRequest("POST", "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res := httptest.NewRecorder()

	uploadHandler(res, req)

	if res.Code != http.StatusOK {
		t.Error("not 200")
	}

	t.Log(res.Body.String())
}
