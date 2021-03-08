package main

import (
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Port string
}

type FileManager struct {
	TempDir string
}

type Config struct {
	Server      Server
	FileManager FileManager
}

func loadEnv(file string, path string, typeFile string) (*Config, error) {
	viper.SetConfigName(file)
	viper.AddConfigPath(path)
	viper.AutomaticEnv()
	viper.SetConfigType(typeFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	port, ok := viper.Get("SERVER.PORT").(string)
	if !ok {
		return nil, fmt.Errorf("invalid type assertion")
	}
	tmpDir, ok := viper.Get("FILE.TMPDIR").(string)

	config := Config{
		Server: Server{
			Port: port,
		},
		FileManager: FileManager{
			TempDir: tmpDir,
		},
	}

	return &config, nil
}
