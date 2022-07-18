package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	SenderInterval uint
	Message        string
}

func NewConfig() *Config {

	senderInterval, err := strconv.ParseUint(os.Getenv("SENDER_INTERVAL"), 10, 64)
	if err != nil {
		log.Fatal("Failed to parse server interval field")
	}

	return &Config{
		SenderInterval: uint(senderInterval),
		Message:        os.Getenv("MESSAGE"),
	}
}
