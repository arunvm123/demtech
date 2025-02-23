package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

var configuration config

type config struct {
	Port     string         `yaml:"port" env:"PORT"`
	Database databaseConfig `yaml:"database" env:"DATABASE"`
}

type databaseConfig struct {
	User         string `yaml:"user" env:"DB_USER"`
	Password     string `yaml:"password" env:"DB_PASSWORD"`
	DatabaseName string `yaml:"database_name" env:"DB_NAME"`
	Host         string `yaml:"host" env:"DB_HOST"`
	Port         string `json:"port" env:"DB_PORT"`
}

func Initialise(filepath string, env bool) (*config, error) {
	var err error

	if env {
		err = cleanenv.ReadEnv(&configuration)
	} else {
		err = cleanenv.ReadConfig(filepath, &configuration)
	}

	if err != nil {
		return nil, err
	}

	return &configuration, nil
}

func GetConfig() *config {
	return &configuration
}
