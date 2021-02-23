package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSettingsPositiveCase(t *testing.T) {
	_, err := loadEnv(confFile, pathConfFile, typeConfFile)
	if err != nil {
		t.Errorf("Error reading config file, %s", err)
	}
}

func TestSettingsNegativeCase(t *testing.T) {
	_, err := loadEnv("confFile", pathConfFile, typeConfFile)
	path, _ := os.Getwd()
	errMessage := fmt.Sprintf("Config File \"confFile\" Not Found in \"[%s]\"", path)
	if fmt.Sprint(err) != errMessage {
		t.Errorf("got: %s, want: %s", err, errMessage)
	}
}
