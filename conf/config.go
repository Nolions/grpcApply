package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Tcp Tcp
}

type Tcp struct {
	Port  string
	Address string
}

var Conf Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func Load() {
	Conf.Tcp = Tcp{
		Port:  os.Getenv("TCP_PORT"),
		Address: os.Getenv("TCP_ADDRESS"),
	}
}
