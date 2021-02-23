package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port string
}

func loadEnv(file string, path string, typeFile string) (*ServerConfig, error) {
	viper.SetConfigName(file)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType(typeFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	port, ok := viper.Get("SERVER.PORT").(string)
	if !ok {
		return nil, fmt.Errorf("Invalid type assertion")
	}
	config := ServerConfig{Port: port}

	return &config, nil
}
