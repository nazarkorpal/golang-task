package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	ServerInterval uint
	Message        string
}

func NewConfig() *Config {

	serverInterval, err := strconv.ParseUint(os.Getenv("SERVER_INTERVAL"), 10, 64)
	if err != nil {
		log.Fatal("Failed to parse server interval field")
	}

	return &Config{
		ServerInterval: uint(serverInterval),
		Message:        os.Getenv("MESSAGE"),
	}
}
