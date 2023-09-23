package app

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type stConfiguration struct {
	ApiAndPort string
}

var config stConfiguration

func LoadConfiguration() error {
	err := godotenv.Load(".env")

	if err != nil {
		err = errors.New("env file not found")
	} else {
		config = stConfiguration{
			ApiAndPort: os.Getenv("ApiAndPort"),
		}
	}
	return err
}

func GetApiAndPort() string {
	return config.ApiAndPort
}
