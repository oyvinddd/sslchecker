package app

import (
	"errors"
	"log"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Address     string
	Environment map[string]string
}

func NewConfig(addr string, envFile *string) Configuration {
	env, err := loadFromEnvironment(envFile)
	if err != nil {
		log.Fatalf("Fatal! Error loading environment variables: %s\n", err.Error())
	}
	return Configuration{
		Address:     addr,
		Environment: env,
	}
}

func (configuration Configuration) EnvironmentVariable(key string) (string, error) {
	value := configuration.Environment[key]
	if value == "" {
		return "", errors.New("no value for key: " + key)
	}
	return value, nil
}

func loadFromEnvironment(filename *string) (map[string]string, error) {
	if filename != nil {
		return godotenv.Read(*filename)
	}
	return godotenv.Read()
}
