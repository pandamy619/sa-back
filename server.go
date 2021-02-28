package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	confFile     = "config"
	pathConfFile = "."
	typeConfFile = "yml"
)

func main() {
	config, err := loadEnv(confFile, pathConfFile, typeConfFile)
	if err != nil {
		fmt.Println(err)
	}
	mux := http.NewServeMux()
	handlers(mux)

	if err := http.ListenAndServe(":"+config.Port, mux); err != nil {
		log.Fatal(err)
	}

}
